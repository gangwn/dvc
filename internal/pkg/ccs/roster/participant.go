package roster

import (
		"github.com/libp2p/go-libp2p-peer"
)

type Participant struct {
	Pid peer.ID
	UserId string
	Name string
}

func NewParticipant(id peer.ID, userId string, name string) *Participant {
	return &Participant{id,userId,name}
}