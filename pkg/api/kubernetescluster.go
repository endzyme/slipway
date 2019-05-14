package api

import (
	"context"

	"github.com/endzyme/slipway/protobuf/slipway"
)

type KubernetesService struct {
	kubeadmPath string
}

func (c KubernetesService) Initialize(request *slipway.InitRequest, resultServer slipway.KubernetesCluster_InitializeServer) error {
	return nil
}

func (c KubernetesService) GetBootstrapToken(ctx context.Context, request *slipway.BootstrapTokenRequest) (*slipway.BootstrapTokenResponse, error) {
	return nil, nil
}

func (c KubernetesService) UploadCertificates(ctx context.Context, request *slipway.UploadCertsRequest) (*slipway.UploadCertsResponse, error) {
	return nil, nil
}
