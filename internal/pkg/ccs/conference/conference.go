package conference

import (
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/libp2p/go-libp2p-peer"
)

type Conference interface {
	JoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest)
	LeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest)
	EndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest)
	Chat(id peer.ID, message *dvc_protocol.ChatMessage)
}