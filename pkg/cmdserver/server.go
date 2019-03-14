package cmdserver

import (
	"context"
	"os/exec"

	"github.com/endzyme/telephone/pkg/gossip"
	"github.com/endzyme/telephone/protobuf/krbrnrtr"
	"github.com/go-chi/chi"
)

type Server struct {
	gossipServer gossip.Server
}

//GetGossipMembers gets the members of the gossip cluster
func (s *Server) GetGossipMembers(ctx context.Context, req *krbrnrtr.GossipMembersRequest) (resp *krbrnrtr.GossipMembersResponse, err error) {
	println(req.String())
	membersList := s.gossipServer.GetGossipMembers()
	resp = &krbrnrtr.GossipMembersResponse{
		Result: membersList,
	}
	return resp, nil
}

// RunAction runs the action
func (s *Server) RunAction(ctx context.Context, req *krbrnrtr.CommandRequest) (*krbrnrtr.CommandResponse, error) {
	cmd := exec.Command(req.Action, req.Arguments...)
	output, err := cmd.CombinedOutput()
	resp := &krbrnrtr.CommandResponse{
		CommandRequestId: "asdf",
		ReceiverNode:     "uhhh wat",
		Status:           krbrnrtr.CommandResponse_ACCEPTED,
		StatusMessage:    string(output),
	}
	return resp, err
}

// Build returns the bind port and http server to listen
func Build(gossipServer gossip.Server) *chi.Mux {
	twirpServer := &Server{
		gossipServer: gossipServer,
	}
	handler := krbrnrtr.NewKrbrnrtrServer(twirpServer, nil)
	routes := chi.NewRouter()
	routes.Mount(krbrnrtr.KrbrnrtrPathPrefix, handler)

	return routes
}
