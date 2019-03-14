package gossip

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/endzyme/telephone/protobuf/krbrnrtr"
	"github.com/hashicorp/serf/serf"
)

//BuildSerfServer returns a build and ready to run gossip server
func BuildSerfServer(eventChannel chan serf.Event, joinAddrs []string) (Server, error) {
	var server SerfGossipServer
	server.joinAddrs = joinAddrs
	err := server.Initialize(eventChannel)
	return &server, err
}

//SerfGossipServer is the interface implementor for a GossipServer
type SerfGossipServer struct {
	config    *serf.Config
	cluster   *serf.Serf
	joinAddrs []string
}

//Stop leaves the cluster and shutsdown the gossip server
func (s SerfGossipServer) Stop() {
	// TODO Should this be checking for errors on these calls
	s.cluster.Leave()
	s.cluster.Shutdown()
}

//Initialize starts the cluster and sets the cluster and config options
func (s *SerfGossipServer) Initialize(eventChannel chan serf.Event) error {
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
	config.EventCh = eventChannel
	config.MemberlistConfig.BindPort, _ = strconv.Atoi(os.Args[1])
	config.NodeName = "XITADO"
	config.TombstoneTimeout = 5 * time.Minute
	config.Tags = map[string]string{
		"Hostname": randomHostname,
		// TODO add more metadata here like cloud availzone and ip or real hostname
		// TODO Add metadata like kubernetes-master or kubernetes-node
	}
	config.ProtocolVersion = serf.ProtocolVersionMax

	// NOTE This is required to accommodate nodes quicking due to connecting with the same name
	// TODO Figure out if there is a better way to shuffle conflict resolution logic
	config.EnableNameConflictResolution = false

	// Setup the cluster with the config
	cluster, err := serf.Create(config)
	if err != nil {
		println(err.Error())
		return err
	}

	// TODO start the cluster and attempt to join any specified joinAddrs
	_, err = cluster.Join(s.joinAddrs, true)
	if err != nil {
		println(err.Error())
	}

	// Set the cluster and config settings for the struct
	s.cluster = cluster
	s.config = config

	return nil
}

//GetEventChannel returns the channel which events can be received on (i think)
func (s SerfGossipServer) GetEventChannel() chan<- serf.Event {
	return s.config.EventCh
}

//GetGossipMembers gets all the members of gossip cluster in the protobuf format
func (s SerfGossipServer) GetGossipMembers() []*krbrnrtr.GossipMember {
	var members []*krbrnrtr.GossipMember
	for _, serfMember := range s.cluster.Members() {
		if serfMember.Status != serf.StatusAlive {
			continue
		}
		gossipMember := &krbrnrtr.GossipMember{}
		gossipMember.Hostname = serfMember.Name
		gossipMember.IpAddress = serfMember.Addr.String()
		for key, val := range serfMember.Tags {
			gossipMember.Labels = append(gossipMember.Labels,
				fmt.Sprintf("%v=%v", key, val))
		}
		members = append(members, gossipMember)
	}
	return members
}
