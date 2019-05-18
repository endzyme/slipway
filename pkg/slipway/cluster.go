package slipway

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/serf/client"
	"github.com/hashicorp/serf/serf"
)

// StartSlipwayCluster returns a build and ready to run gossip server
func StartSlipwayCluster(eventChannel chan serf.Event, config *SlipwayConfig) (SlipwayCluster, error) {
	server := cluster{
		eventChannel: eventChannel,
		bindPort:     config.GossipBindPort,
		secretKey:    config.GossipSecret,
	}
	err := server.Start()
	if err != nil {
		return &server, err
	}

	go server.listenForUserEvents()

	server.Join(config.GossipJoinAddrs...)

	return &server, err
}

//cluster is the interface implementor for a GossipServer
type cluster struct {
	bindPort     int
	eventChannel chan serf.Event
	secretKey    []byte
	config       *serf.Config
	cluster      *serf.Serf
	joinAddrs    []string
	clientConfig *client.Config
	sync.Mutex
}

//Stop leaves the cluster and shutsdown the gossip server
func (c *cluster) Stop() {
	c.Lock()
	defer c.Unlock()
	// TODO Should this be checking for errors on these calls
	c.cluster.Leave()
	c.cluster.Shutdown()
}

//Start starts the cluster and sets the cluster and config options
func (c *cluster) Start() error {
	randomHostname := func(n int) string {
		bytes := make([]byte, n)
		for i := 0; i < n; i++ {
			bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
		}
		return string(bytes)
	}(6)

	// Initialize the config
	config := serf.DefaultConfig()
	config.Init()
	config.CoalescePeriod = time.Millisecond * 100
	config.MemberlistConfig.SecretKey = c.secretKey
	config.EventCh = c.eventChannel
	config.MemberlistConfig.BindPort = c.bindPort
	config.NodeName = randomHostname
	config.TombstoneTimeout = 5 * time.Minute
	config.Tags = map[string]string{
		"Hostname": randomHostname,
		// TODO add more metadata here like cloud availzone and ip or real hostname
		// TODO Add metadata like kubernetes-master or kubernetes-node
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
	c.cluster = cluster
	c.config = config

	return nil
}

// Join allows you to join cluster nodes with a list of addrs:ports
func (c *cluster) Join(addrs ...string) error {
	if len(addrs) > 0 {
		c.Lock()
		defer c.Unlock()

		// attempt to join any specified joinAddrs
		log.Printf("Trying to join (%v)", addrs)
		_, err := c.cluster.Join(addrs, true)
		if err != nil {
			return err
		}
	} else {
		log.Print("No join addrs defined, waiting for peers...")
	}

	return nil
}

func (c *cluster) listenForUserEvents() {
	for {
		select {
		case event := <-c.eventChannel:
			// TODO make this a real thing
			if event.EventType().String() == "user" {
				fmt.Printf("%v\n", c.config.Tags)
				message := strings.SplitN(event.String(), ": ", 2)
				c.AddTag("last-user-event", message[1])
				fmt.Printf("%v: %v\n", event.EventType(), event.String())
				fmt.Printf("%v\n", c.config.Tags)
			}
		}
	}
}

// BroadcastEvent sends events to the gossip cluster
func (c *cluster) BroadcastEvent(msg string) error {
	return c.cluster.UserEvent(msg, []byte(msg), false)
}

// AddTag creates or updates a tag on the cluster
func (c *cluster) AddTag(key, value string) error {
	log.Print(c.config.Tags)
	tags := c.config.Tags
	tags[key] = value

	return c.cluster.SetTags(tags)
}

// RemoveTag takes a tag off the cluster. Safe to call on tags that don't exist
func (c *cluster) RemoveTag(key string) error {
	c.Lock()
	defer c.Unlock()

	tags := make(map[string]string)
	for k, v := range c.config.Tags {
		if key != k {
			tags[key] = v
		}
	}

	return c.cluster.SetTags(tags)
}

// WatchNodeStatus starts a loop to transition the readiness state of the
// cluster between lifecycle states.
func (c *cluster) WatchNodeStatus() {

}
