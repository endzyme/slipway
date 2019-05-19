package main

import (
	"fmt"
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

	// initialize cluster instance
	ch := make(chan serf.Event)
	slipwayCluster, err := slipway.StartSlipwayCluster(ch, config)
	if err != nil {
		log.Fatalf("Error Building Gossip Cluster: (%v)", err)
	}
	defer slipwayCluster.Stop()

	var listMembersUsr1 = make(chan os.Signal)
	signal.Notify(listMembersUsr1, syscall.SIGUSR1)
	go func() {
		for {
			select {
			case <-listMembersUsr1:
				members, err := slipwayCluster.GetMembers(nil, nil)
				fmt.Printf("Members: (%v)\nError: (%v)\n", members, err)
			}
		}
	}()

	// Start up a signal channel for graceful termination
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGHUP)

	// start the api server and await commands
	if err = api.ServeGRPC(config.APIBind, slipwayCluster, gracefulStop); err != nil {
		log.Fatal(err)
	}
}
