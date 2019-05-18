package slipway

import (
	"log"
	"os"
	"strconv"
)

var (
	// slipway server node config
	nodeRole = "master"

	// slipway API interface config
	apiBind = "0.0.0.0:7070"

	// slipway gossip cluster config
	gossipBindPort  = 7946
	gossipSecret    = "30a4817deb1fce06fa1e3452445a9742"
	gossipJoinAddrs = []string{}
)

// TODO add secret key and role type for tags

func ReadConfigs() *SlipwayConfig {
	var err error
	if len(os.Args) > 3 {
		gossipJoinAddrs = []string{os.Args[3]}
	}

	if len(os.Args) > 2 {
		apiBind = os.Args[2]
	}

	if len(os.Args) > 1 {
		gossipBindPort, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Could not convert %v into integer.", os.Args[1])
		}
	}

	config := SlipwayConfig{
		NodeRole:        nodeRole,
		APIBind:         apiBind,
		GossipBindPort:  gossipBindPort,
		GossipJoinAddrs: gossipJoinAddrs,
		GossipSecret:    []byte(gossipSecret),
	}

	return &config
}

// The slipway configuration
type SlipwayConfig struct {
	NodeRole string

	APIBind string

	GossipBindPort  int
	GossipJoinAddrs []string
	GossipSecret    []byte
}
