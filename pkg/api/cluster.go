package api

import (
	"context"

	"github.com/endzyme/slipway/pkg/slipway"
	"github.com/endzyme/slipway/protobuf/slipway"
)

type ClusterHandler struct {
	slipwayCluster slipway.SlipwayCluster
}

func (c ClusterHandler) ListMembers(ctx context.Context, request *slipway.MembersRequest) (*slipway.MembersResponse, error) {
	return nil, nil
}
