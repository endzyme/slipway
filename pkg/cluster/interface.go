package cluster

// SlipwayClusterServer Interface is for starting and stoping the server
type SlipwayClusterServer interface {
	Start() error
	Stop()
}
