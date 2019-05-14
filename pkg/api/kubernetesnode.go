package api

import (
	"context"

	"github.com/endzyme/slipway/protobuf/slipway"
)

type KubernetesNodeHandler struct{}

func (k KubernetesNodeHandler) GetBootstrapToken(ctx context.Context, request *slipway.BootstrapTokenRequest) (*slipway.BootstrapTokenResponse, error) {
	return nil, nil
}

func (k KubernetesNodeHandler) BootstrapMaster(ctx context.Context, request *slipway.BootstrapMasterRequest) (*slipway.BootstrapTokenResponse, error) {
	return nil, nil
}
