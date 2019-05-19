package slipway

import "net"

// Cluster Interface is for starting and stoping the server
type Cluster interface {
	Start() error
	Stop()
	Join(addr ...string) (int, error)
	BroadcastEvent(string) error
	AddTag(key, value string) error
	RemoveTag(key string) error
	GetMembers(tagsFilter map[string]string, membersStatusFilter []GossipMemberStatus) ([]ClusterMember, error)
}

// ClusterMember describes a single cluster member
type ClusterMember struct {
	ID             string
	Hostname       string
	IPAddress      net.IP
	Tags           map[string]string
	LifecycleState NodeState
	MemberStatus   GossipMemberStatus
}

// GossipMemberStatus is a status for gossip members
type GossipMemberStatus int

func (g GossipMemberStatus) String() string {
	switch g {
	case GossipStatusAvailable:
		return "available"
	case GossipStatusUnavailable:
		return "unavailable"
	case GossipStatusUnknown:
		return "unknown"
	default:
		panic("Broken GossipMemberStatus. Cannot find expected set of constants.")
	}
}

const (
	// GossipStatusUnknown to catch new serf statuses
	GossipStatusUnknown GossipMemberStatus = iota
	// GossipStatusAvailable for status available
	GossipStatusAvailable
	// GossipStatusUnavailable to summarize non healthy cluster node statuses
	GossipStatusUnavailable
)
