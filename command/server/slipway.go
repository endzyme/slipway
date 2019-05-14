package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/braintree/manners"
	"github.com/hashicorp/serf/serf"

	"github.com/endzyme/slipway/pkg/api"
	"github.com/endzyme/slipway/pkg/gossip"
)

func main() {
	defer println("exiting cleanly")

	var webPort string
	var gossipPort int
	var gossipJoinAddrs []string

	if len(os.Args) > 3 {
		gossipJoinAddrs = []string{os.Args[3]}
	} else {
		gossipJoinAddrs = []string{""}
	}

	if len(os.Args) > 2 {
		webPort = os.Args[2]
	} else {
		webPort = "0.0.0.0:7070"
	}

	if len(os.Args) > 1 {
		var err error
		gossipPort, err = strconv.Atoi(os.Args[1])
		if err != nil {
			gossipPort = 7946
		}
	} else {
		gossipPort = 7946
	}

	log.Printf("Web Port: %v\n", webPort)
	log.Printf("Gossip Port: %v\n", gossipPort)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 second to finish processing")
		manners.Close()
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	// read configuration
	// initialize cluster instance
	// initialize http server
	// sit and wait for signals to pass to downstream components
	ch := make(chan serf.Event)
	gossipServer, err := gossip.BuildSerfServer(ch, gossipPort, gossipJoinAddrs)
	if err != nil {
		println("FAILED BUILDING GOSSIP SERVER")
		os.Exit(1)
	}
	defer gossipServer.Stop()

	go func() {
		for {
			select {
			case event := <-ch:
				fmt.Println("Hit")
				fmt.Println(event.String())
			}

		}
	}()

	if err = api.Build(webPort, gossipServer); err != nil {
		log.Fatal(err)
	}
}
