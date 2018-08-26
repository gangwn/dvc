package conference

import (
		"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"github.com/gangwn/dvc/internal/pkg/ccs"
)

type Conference interface {
	ConferenceId() (string)
	Roster() (*roster.Roster)
	JoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest) ccs.CCSErrorCode
	LeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest) ccs.CCSErrorCode
	EndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest) ccs.CCSErrorCode
	Chat(id peer.ID, message *dvc_protocol.ChatMessage)
}