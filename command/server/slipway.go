package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/serf/serf"

	"github.com/endzyme/slipway/pkg/api"
	"github.com/endzyme/slipway/pkg/cluster"
)

func main() {
	defer println("exiting cleanly")

	// read configurations

	// connect to gossip to join slipway cluster or await connections for 2 hours (then die.)

	// start the api server and await commands

	log.Printf("Web Port: %v\n", apiBind)
	log.Printf("Gossip Port: %v\n", gossipBindPort)
	log.Printf("Gossip Port: %v\n", gossipJoinAddrs)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

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
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		fmt.Println("Wait for 2 second to finish processing")
		slipwayCluster.Stop()
		time.Sleep(2 * time.Second)
	}()

	go func() {
		for {
			select {
			case event := <-ch:
				fmt.Println("Hit")
				fmt.Println(event.String())
			}

		}
	}()

	if err = api.ServeGRPC(apiBind, slipwayCluster, gracefulStop); err != nil {
		log.Fatal(err)
	}
}
