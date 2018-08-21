package roster

import (
	"github.com/gangwn/dvc/pkg/net"
	"github.com/libp2p/go-libp2p-peer"
)

type Participant struct {
	Id peer.ID
	name string
	Node *net.Node
}

func NewParticipant(id peer.ID, name string, node *net.Node) *Participant {
	return &Participant{id,name,node}
}