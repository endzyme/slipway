package cluster

import (
	"math/rand"
	"time"

	"github.com/hashicorp/serf/serf"
)

//BuildSerfServer returns a build and ready to run gossip server
func StartSlipwayCluster(eventChannel chan serf.Event, bindPort int, joinAddrs []string, secretKey []byte) (SlipwayClusterServer, error) {
	server := SerfGossipServer{
		joinAddrs:    joinAddrs,
		eventChannel: eventChannel,
		bindPort:     bindPort,
		secretKey:    secretKey,
	}
	err := server.Start()
	return &server, err
}

//SerfGossipServer is the interface implementor for a GossipServer
type SerfGossipServer struct {
	bindPort     int
	eventChannel chan serf.Event
	secretKey    []byte
	config       *serf.Config
	cluster      *serf.Serf
	joinAddrs    []string
}

//Stop leaves the cluster and shutsdown the gossip server
func (s SerfGossipServer) Stop() {
	// TODO Should this be checking for errors on these calls
	s.cluster.Leave()
	s.cluster.Shutdown()
}

//Start starts the cluster and sets the cluster and config options
func (s *SerfGossipServer) Start() error {
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
