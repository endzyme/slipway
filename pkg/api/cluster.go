package api

import (
	"context"

	"github.com/endzyme/slipway/pkg/slipway"
	pb "github.com/endzyme/slipway/protobuf/slipway"
)

type ClusterHandler struct {
	slipwayCluster slipway.SlipwayCluster
}

func (c ClusterHandler) ListMembers(ctx context.Context, request *pb.MembersRequest) (*pb.MembersResponse, error) {
	return nil, nil
}
