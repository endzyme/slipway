package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/endzyme/slipway/pkg/api"
	"github.com/endzyme/slipway/pkg/slipway"
	"github.com/hashicorp/serf/serf"
)

func main() {
	defer println("exiting cleanly")

	// read configurations
	// var reloadChannel = make(chan os.Signal)
	// signal.Notify(reloadChannel, syscall.SIGUSR1)
	config := slipway.ReadConfigs()

	// Start up a signal channel for graceful termination
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGHUP)

	// initialize cluster instance
	ch := make(chan serf.Event)
	slipwayCluster, err := slipway.StartSlipwayCluster(ch, config)
	if err != nil {
		println("FAILED BUILDING GOSSIP SERVER")
		os.Exit(1)
	}
	defer slipwayCluster.Stop()

	// run through your start up sequence and continually scan for state with which to  update tags in gossip
	// slipway.WatchNodeStatus()

	// start the api server and await commands
	if err = api.ServeGRPC(config.APIBind, slipwayCluster, gracefulStop); err != nil {
		log.Fatal(err)
	}
}
