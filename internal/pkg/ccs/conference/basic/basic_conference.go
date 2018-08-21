package basic

import (
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/pkg/protocol"
)

type BasicConference struct {
	confId string
	roster *roster.Roster
}

func (conf *BasicConference) JoinConference(message dvc_protocol.JoinConferenceMessage) {

	//participant := roster.NewParticipant(id, join_conf_msg.GetName(), nil)
}