package slipway

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/serf/client"
	"github.com/hashicorp/serf/serf"
)

// StartSlipwayCluster returns a build and ready to run gossip server
func StartSlipwayCluster(eventChannel chan serf.Event, config *Config) (Cluster, error) {
	server := cluster{
		eventChannel:    eventChannel,
		bindPort:        config.GossipBindPort,
		secretKey:       config.GossipSecret,
		initialJoinList: config.GossipJoinAddrs,
	}
	err := server.Start()
	if err != nil {
		return &server, err
	}

	// Start the node status watcher
	go server.watchNodeLifecycle()

	// Try to join any other cluster members
	if err := server.startInitialJoin(server.initialJoinList); err != nil {
		return &server, err
	}

	// start the Event Listener
	go server.listenForUserEvents()

	return &server, err
}

//cluster is the interface implementor for a GossipServer
type cluster struct {
	bindPort        int
	eventChannel    chan serf.Event
	secretKey       []byte
	serfConfig      *serf.Config
	serf            *serf.Serf
	initialJoinList []string
	clientConfig    *client.Config
	sync.Mutex
}

//Stop leaves the cluster and shutsdown the gossip server
func (c *cluster) Stop() {
	c.Lock()
	defer c.Unlock()
	// TODO Should this be checking for errors on these calls
	c.serf.Leave()
	c.serf.Shutdown()
}

//Start starts the cluster and sets the cluster and config options
func (c *cluster) Start() error {
	randomID := func(n int) string {
		bytes := make([]byte, n)
		for i := 0; i < n; i++ {
			bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
		}
		return string(bytes)
	}(6)

	// Initialize the config
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Can't get hostname")
	}

	config := serf.DefaultConfig()
	config.Init()
	config.CoalescePeriod = time.Millisecond * 100
	config.MemberlistConfig.SecretKey = c.secretKey
	config.EventCh = c.eventChannel
	config.MemberlistConfig.BindPort = c.bindPort
	config.NodeName = fmt.Sprintf("%v-%v", hostname, randomID)
	config.ReconnectTimeout = 3 * time.Minute
	config.TombstoneTimeout = 5 * time.Minute
	config.Tags = map[string]string{
		"Hostname": hostname,
	}
	config.ProtocolVersion = serf.ProtocolVersionMax

	// NOTE This is required to accommodate nodes quiting due to connecting with the same name
	// TODO Figure out if there is a better way to shuffle conflict resolution logic
	config.EnableNameConflictResolution = false

	// Setup the cluster with the config and start the cluster
	cluster, err := serf.Create(config)
	if err != nil {
		println(err.Error())
		return err
	}

	// Set the cluster and config settings for the struct
	c.serf = cluster
	c.serfConfig = config

	return nil
}

// Join allows you to join cluster nodes with a list of addrs:ports
func (c *cluster) Join(addrs ...string) (int, error) {
	if len(addrs) > 0 {
		c.Lock()
		defer c.Unlock()

		// attempt to join any specified joinAddrs
		contactedNodes, err := c.serf.Join(addrs, false)
		if err != nil {
			return contactedNodes, err
		}
		return contactedNodes, nil
	}

	log.Print("No join addrs defined, waiting for peers...")

	return 0, nil
}

func (c *cluster) listenForUserEvents() {
	for {
		select {
		case event := <-c.eventChannel:
			// TODO make this a real thing
			if event.EventType().String() == "user" {
				fmt.Printf("%v\n", c.serfConfig.Tags)
				message := strings.SplitN(event.String(), ": ", 2)
				c.AddTag("last-user-event", message[1])
				fmt.Printf("%v: %v\n", event.EventType(), event.String())
				fmt.Printf("%v\n", c.serfConfig.Tags)
			}
		}
	}
}

// BroadcastEvent sends events to the gossip cluster
func (c *cluster) BroadcastEvent(msg string) error {
	return c.serf.UserEvent(msg, []byte(msg), false)
}

// AddTag creates or updates a tag on the cluster
func (c *cluster) AddTag(key, value string) error {
	tags := c.serfConfig.Tags
	tags[key] = value

	return c.serf.SetTags(tags)
}

// RemoveTag takes a tag off the cluster. Safe to call on tags that don't exist
func (c *cluster) RemoveTag(key string) error {
	c.Lock()
	defer c.Unlock()

	tags := make(map[string]string)
	for k, v := range c.serfConfig.Tags {
		if key != k {
			tags[key] = v
		}
	}

	return c.serf.SetTags(tags)
}

// watchNodeLifecycle starts a loop to transition the readiness state of the
// cluster between lifecycle states.
func (c *cluster) watchNodeLifecycle() {
	// initially set the first state to waiting to bootstrap
	c.AddTag("LifecycleState", NodeWaitingToBootstrap.String())

	// watch a channel for state changes from node package to perform some cluster tag changes
	// nodeStateChannel := node.WatchNodeState()
	// for {
	// 	select newState := <-nodeStateChannel {
	// 	case NodeReady:
	// 		// do stuff
	// 	case NodeFailedBootstrap:
	// 		// do stuff
	// 	case NodeWaitingToBootstrap:
	// 		// do stuff
	// 	case NodeBootstrapping:
	// 		// do stuff
	// 	}
	// }
	//
}

// GetMembers returns members of gossip cluster
// Requires all tags AND *any* of the listed statuses
// TODO this needs testing for filter bugs
func (c *cluster) GetMembers(filterTags map[string]string, membersStatusFilter []GossipMemberStatus) ([]ClusterMember, error) {
	var members []ClusterMember

	serfMembers := c.serf.Members()

MemberLoop:
	for _, serfMember := range serfMembers {
		// filter out any members that aren't of a desired memberStatus
		if len(membersStatusFilter) > 0 {
			statusFound := false
			for _, expectedStatus := range membersStatusFilter {
				gossipStatus := translateSerfStatusToGossipStatus(serfMember.Status)

				// set found to true and break from loop of status filters
				if gossipStatus == expectedStatus {
					statusFound = true
					break
				}
			}

			// skip this member if statuses were not found
			if !statusFound {
				continue MemberLoop
			}
		}

		// filter the members with filterTags if any are defined
		if len(filterTags) > 0 {
			// if the serfMember is missing any of the tags then skip this
			// member from the loop and break early if even one tag is missing
			// or if the tags value doesn't match disired tag value.
			for filterTagKey, filterTagValue := range filterTags {
				var serfMemberTagValue string
				var keyFound bool

				// check to see if serfMember.Tags has the filterTagKey
				if serfMemberTagValue, keyFound = serfMember.Tags[filterTagKey]; !keyFound {
					// Member was missing on of the tag keys, skip from response
					continue MemberLoop
				}

				// validate the serfMember.Tag Value == filterTagValue
				if filterTagValue != serfMemberTagValue {
					// serfMembers tag key value didn't match disired filter value, skip this member from response
					continue MemberLoop
				}
			}
		}

		clusterMember, err := translateSerfMemberToClusterMember(serfMember)
		if err != nil {
			log.Printf("Error translating serfMember to clusterMember struct: %v", err)
		}

		members = append(members, clusterMember)
	}

	if len(members) == 0 {
		return members, fmt.Errorf("No Members Found using Tags: (%v)", filterTags)
	}

	return members, nil
}

func (c *cluster) GetReadyMembers() ([]ClusterMember, error) {
	return c.GetMembers(map[string]string{"NodeState": "Ready"}, nil)
}

func translateSerfMemberToClusterMember(serfMember serf.Member) (ClusterMember, error) {
	m := ClusterMember{
		ID:           serfMember.Name,
		IPAddress:    serfMember.Addr,
		Tags:         serfMember.Tags,
		MemberStatus: translateSerfStatusToGossipStatus(serfMember.Status),
	}

	if hostname, found := serfMember.Tags["Hostname"]; found {
		m.Hostname = hostname
	} else {
		return m, fmt.Errorf("Hostname Tag Not Found in Tags! Tags:(%v)", serfMember.Tags)
	}

	if _, found := serfMember.Tags["LifecycleState"]; found {
		// if lifecycleState, found := serfMember.Tags["LifecycleState"]; found {
		// m.LifecycleState = lifecycleState
	} else {
		return m, fmt.Errorf("LifecycleStatus Tag Not Found in Tags! Tags:(%v)", serfMember.Tags)
	}

	return m, nil
}

func translateSerfStatusToGossipStatus(status serf.MemberStatus) GossipMemberStatus {
	var memberStatus GossipMemberStatus

	// if we need a specific status
	switch status {
	case serf.StatusNone, serf.StatusFailed, serf.StatusLeaving, serf.StatusLeft:
		memberStatus = GossipStatusUnavailable
	case serf.StatusAlive:
		memberStatus = GossipStatusAvailable
	default:
		memberStatus = GossipStatusUnknown
	}

	return memberStatus
}

// startInitialJoin will try to connect
func (c *cluster) startInitialJoin(joinList []string) error {
	retryInterval := time.Tick(time.Second * 1)
	timeoutPeriod := time.Second * 30
	timeout := time.After(timeoutPeriod)

	log.Printf("Attempting to join cluster members (%v) for (%v)...", joinList, timeoutPeriod)
	for {
		select {
		case <-retryInterval:
			contactedNodes, err := c.Join(joinList...)
			if err != nil {
				if contactedNodes == 0 {
					continue
				}
			}
			return nil
			// case <-time.After(timeoutPeriod):
		case <-timeout:
			return fmt.Errorf("Could not connect to cluster within timeout period (%v)", timeoutPeriod)
		}
	}
}
