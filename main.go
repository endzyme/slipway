package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"

	"github.com/endzyme/telephone/pkg/cmdserver"
	"github.com/endzyme/telephone/pkg/gossip"
	"github.com/hashicorp/serf/serf"
)

func main() {
	defer println("exiting cleanly")

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
	gossipServer, err := gossip.BuildSerfServer(ch, []string{"127.0.0.1:7947"})
	if err != nil {
		println("FAILED BUILDING GOSSIP SERVER")
		os.Exit(1)
	}
	defer gossipServer.Stop()

	// channel := gossipServer.GetEventChannel()

	webServer := cmdserver.Build(gossipServer)
	// manners.ListenAndServe(":7070", webServer)
	manners.ListenAndServe(os.Args[2], webServer)
}
