package kubeadm

import (
	"fmt"
	"time"

	command "github.com/go-cmd/cmd"
)

// Client is a kubeadm client
//   NOTE: there is a hard timeout on all commands to kubeadm that only allow
//         commands to run in the background for 60 minutes.
type Client struct {
	KubeadmPath     string
	hardKillTimeout time.Duration
}

// NewClient creates a compliant new kubeadm client
func NewClient() (Client, error) {
	// try to find client in path etc
	client := Client{
		KubeadmPath:     "kubeadm",
		hardKillTimeout: time.Minute * 60,
	}

	// try to run command and output the version we initialized
	if run, err := client.RunSynchronously("help"); err != nil {
		return client, fmt.Errorf("Error processing command: %v", run.Error)
	}

	// return
	return client, nil
}

// Run the command in the background with a channel for status and a timeout
func (c Client) Run(subcmd string, args ...string) (*command.Cmd, <-chan command.Status) {
	commandOpts := command.Options{
		Buffered:  false,
		Streaming: true,
	}
	commandArguments := []string{subcmd}

	for _, arg := range args {
		commandArguments = append(commandArguments, arg)
	}

	cmd := command.NewCmdOptions(commandOpts, c.KubeadmPath, commandArguments...)
	resp := cmd.Start()

	// stop command if exceeds timeout duration
	go func() {
		<-time.After(c.hardKillTimeout)
		cmd.Stop()
	}()

	return cmd, resp
}

// RunSynchronously runs the kubeadm command inline (with a timeout)
func (c Client) RunSynchronously(subcmd string, args ...string) (command.Status, error) {
	commandOpts := command.Options{
		Buffered:  true,
		Streaming: false,
	}
	commandArguments := []string{subcmd}

	for _, arg := range args {
		commandArguments = append(commandArguments, arg)
	}

	run := command.NewCmdOptions(commandOpts, c.KubeadmPath, commandArguments...)
	cmd := run.Start()

	// stop command if exceeds timeout duration
	go func() {
		<-time.After(c.hardKillTimeout)
		run.Stop()
	}()

	resp := <-cmd
	if !resp.Complete {
		return resp, fmt.Errorf("Failed to complete command: %v\nStdout: %v\nStderr: %v", resp.Error, resp.Stdout, resp.Stderr)
	}

	return resp, nil
}
