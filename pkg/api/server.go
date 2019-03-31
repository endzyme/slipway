package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"sync"
	"time"

	"github.com/go-cmd/cmd"

	"github.com/endzyme/telephone/pkg/gossip"
	"github.com/endzyme/telephone/protobuf/krbrnrtr"
	"google.golang.org/grpc"
)

type Server struct {
	gossipServer gossip.Server
	mutex        sync.Mutex
}

//GetGossipMembers gets the members of the gossip cluster
func (s Server) GetGossipMembers(ctx context.Context, req *krbrnrtr.GossipMembersRequest) (resp *krbrnrtr.GossipMembersResponse, err error) {
	eventStream := s.gossipServer.SendEvent("fudge", []byte("uhhhhhh"))
	if err != nil {
		println(eventStream.Error())
	}

	membersList := s.gossipServer.GetGossipMembers()
	resp = &krbrnrtr.GossipMembersResponse{
		Result: membersList,
	}
	return resp, nil
}

// RunAction runs the action
func (s Server) RunAction(ctx context.Context, req *krbrnrtr.CommandRequest) (*krbrnrtr.CommandResponse, error) {
	cmd := exec.Command(req.Action, req.Arguments...)
	output, err := cmd.CombinedOutput()
	resp := &krbrnrtr.CommandResponse{
		CommandRequestId: "asdf",
		ReceiverNode:     "uhhh wat",
		Status:           krbrnrtr.CommandResponse_ACCEPTED,
		StatusMessage:    output,
	}
	return resp, err
}

//RunActionSyncResult asdf
func (s Server) RunActionSyncResult(req *krbrnrtr.CommandRequest, resultServer krbrnrtr.Krbrnrtr_RunActionSyncResultServer) error {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	action := cmd.NewCmdOptions(cmdOptions, req.Action, req.Arguments...)
	receiver := "self"

	go func() {
		var msg krbrnrtr.CommandResponse
		for {
			select {
			case line := <-action.Stdout:
				msg.ReceiverNode = receiver
				msg.StatusMessage = []byte(fmt.Sprintf("%s: %s", time.Now().String(), line))
			case line := <-action.Stderr:
				msg.ReceiverNode = receiver
				msg.StatusMessage = []byte(line)
			}
			resultServer.Send(&msg)
		}
	}()

	<-action.Start()

	for len(action.Stdout) > 0 || len(action.Stderr) > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	log.Printf("Finished command %v", action.Name)

	if action.Status().Exit > 0 {
		return fmt.Errorf("Error found executing thing. Exit Code: %d", action.Status().Exit)
	}

	return nil
}

// Build returns the bind port and http server to listen
func Build(bindAddress string, gossipServer gossip.Server) error {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return err
	}

	server := Server{
		gossipServer: gossipServer,
	}

	krbrnrtr.RegisterKrbrnrtrServer(grpcServer, server)
	return grpcServer.Serve(listener)
}
