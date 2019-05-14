package gossip

import (
	"github.com/endzyme/slipway/protobuf/krbrnrtr"
	"github.com/hashicorp/serf/serf"
)

//Server facilitates starting stopping and init of gossip server cluster
type Server interface {
	Initialize(bindPort int, eventChannel chan serf.Event) error
	Stop()
	GetEventChannel() chan<- serf.Event
	GetGossipMembers() []*krbrnrtr.GossipMember
	SendEvent(name string, payload []byte) error
}
