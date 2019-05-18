package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/endzyme/slipway/pkg/api"
	"github.com/endzyme/slipway/pkg/cluster"
	"github.com/hashicorp/serf/serf"
)

func main() {
	defer println("exiting cleanly")

	// read configurations
	// var reloadChannel = make(chan os.Signal)
	// signal.Notify(reloadChannel, syscall.SIGUSR1)
	readConfigs()

	// Note the configs
	log.Printf("Web Port: %v\n", apiBind)
	log.Printf("Gossip Port: %v\n", gossipBindPort)
	log.Printf("Gossip Port: %v\n", gossipJoinAddrs)

	// Start up a signal channel for graceful termination
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGHUP)

	// read configuration
	// initialize cluster instance
	// initialize http server
	// sit and wait for signals to pass to downstream components
	ch := make(chan serf.Event)
	slipwayCluster, err := cluster.StartSlipwayCluster(ch, gossipBindPort, gossipJoinAddrs, []byte(gossipSecret))
	if err != nil {
		println("FAILED BUILDING GOSSIP SERVER")
		os.Exit(1)
	}
	defer slipwayCluster.Stop()

	slipwayCluster.Join(gossipJoinAddrs...)

	// run through your start up sequence and continually scan for state with which to  update tags in gossip
	// cluster.ScanForState()

	// start the api server and await commands
	if err = api.ServeGRPC(apiBind, slipwayCluster, gracefulStop); err != nil {
		log.Fatal(err)
	}
}
