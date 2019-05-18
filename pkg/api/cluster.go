package api

import (
	"context"

	"github.com/endzyme/slipway/pkg/cluster"
	"github.com/endzyme/slipway/protobuf/slipway"
)

type ClusterHandler struct {
	slipwayCluster cluster.SlipwayCluster
}

func (c ClusterHandler) ListMembers(ctx context.Context, request *slipway.MembersRequest) (*slipway.MembersResponse, error) {
	return nil, nil
}
