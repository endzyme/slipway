package kubeadm

import "github.com/go-cmd/cmd"

type KubeadmInterface interface {
	Run(subcmd string, args ...string) (*cmd.Cmd, <-chan cmd.Status)
	RunSynchronously(subcmd string, args ...string) (cmd.Status, error)
}
