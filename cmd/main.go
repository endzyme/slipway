package main

import (
	"log"
	"time"

	"github.com/endzyme/slipway/pkg/kubeadm"
)

func main() {
	var client kubeadm.KubeadmInterface
	client = kubeadm.Client{
		KubeadmPath: "ping",
	}

	command, statusChannel := client.Run("asdfajfihoifhahflafds.cm", "-c", "2")
	// response, err := client.RunSynchronously("gohgsdfghj.cm", "-c", "6")
	// if err != nil {
	// 	log.Fatal(response.Error)
	// }

	go func() {
		duration := time.Second * 7
		<-time.After(duration)
		log.Printf("Failed due to timeout: %v\n", duration)
		command.Stop()
	}()

	go func() {
		for {
			select {
			case line := <-command.Stdout:
				log.Print(line)
			case line := <-command.Stderr:
				log.Print(line)
			}
		}
	}()

	response := <-statusChannel

	log.Print(response.Complete)
	log.Print(response.Exit)
	log.Print(response.Runtime)
	log.Print(response.Error)
	for _, line := range response.Stdout {
		log.Print(line)
	}
	for _, line := range response.Stderr {
		log.Print(line)
	}
}
