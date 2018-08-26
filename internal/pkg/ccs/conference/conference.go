package conference

import (
		"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/pkg/protocol/pb"
)

type Conference interface {
	ConferenceId() (string)
	Roster() (*roster.Roster)
	JoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest)
	LeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest)
	EndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest)
	Chat(id peer.ID, message *dvc_protocol.ChatMessage)
}