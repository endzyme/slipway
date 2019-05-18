package cluster

// SlipwayCluster Interface is for starting and stoping the server
type SlipwayCluster interface {
	Start() error
	Stop()
	Join(addr ...string) error
	SendEvent(string) error
	AddTag(key, value string) error
	RemoveTag(key string) error
}
