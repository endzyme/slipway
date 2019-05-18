package cluster

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
func StartSlipwayCluster(eventChannel chan serf.Event, bindPort int, joinAddrs []string, secretKey []byte) (SlipwayCluster, error) {
	server := SerfCluster{
		eventChannel: eventChannel,
		bindPort:     bindPort,
		secretKey:    secretKey,
	}
	err := server.Start()
	if err != nil {
		return &server, err
	}

	go server.listenForUserEvents()

	return &server, err
}

//SerfCluster is the interface implementor for a GossipServer
type SerfCluster struct {
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
func (s *SerfCluster) Stop() {
	s.Lock()
	defer s.Unlock()
	// TODO Should this be checking for errors on these calls
	s.cluster.Leave()
	s.cluster.Shutdown()
}

//Start starts the cluster and sets the cluster and config options
func (s *SerfCluster) Start() error {
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
	config.MemberlistConfig.SecretKey = s.secretKey
	config.EventCh = s.eventChannel
	config.MemberlistConfig.BindPort = s.bindPort
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
	s.cluster = cluster
	s.config = config

	return nil
}

// Join allows you to join cluster nodes with a list of addrs:ports
func (s *SerfCluster) Join(addrs ...string) error {
	if len(addrs) > 0 {
		s.Lock()
		defer s.Unlock()

		// attempt to join any specified joinAddrs
		log.Printf("Trying to join (%v)", addrs)
		_, err := s.cluster.Join(addrs, true)
		if err != nil {
			return err
		}
	} else {
		log.Print("No join addrs defined, waiting for peers...")
	}

	return nil
}

func (s *SerfCluster) listenForUserEvents() {
	for {
		select {
		case event := <-s.eventChannel:
			// TODO make this a real thing
			if event.EventType().String() == "user" {
				fmt.Printf("%v\n", s.config.Tags)
				message := strings.SplitN(event.String(), ": ", 2)
				s.AddTag("last-user-event", message[1])
				fmt.Printf("%v: %v\n", event.EventType(), event.String())
				fmt.Printf("%v\n", s.config.Tags)
			}
		}
	}
}

// SendEvent sends events to the gossip cluster
func (s *SerfCluster) SendEvent(msg string) error {
	return s.cluster.UserEvent(msg, []byte(msg), false)
}

// AddTag creates or updates a tag on the cluster
func (s *SerfCluster) AddTag(key, value string) error {
	log.Print(s.config.Tags)
	tags := s.config.Tags
	tags[key] = value

	return s.cluster.SetTags(tags)
}

// RemoveTag takes a tag off the cluster. Safe to call on tags that don't exist
func (s *SerfCluster) RemoveTag(key string) error {
	s.Lock()
	defer s.Unlock()

	tags := make(map[string]string)
	for k, v := range s.config.Tags {
		if key != k {
			tags[key] = v
		}
	}

	return s.cluster.SetTags(tags)
}
