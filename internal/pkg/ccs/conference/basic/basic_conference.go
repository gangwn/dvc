package basicconference

import (
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/net"
	"github.com/gangwn/dvc/internal/pkg/ccs/handler"
	"github.com/gangwn/dvc/internal/pkg/ccs"
	"github.com/gangwn/dvc/pkg/protocol/pb"
)

type BasicConference struct {
	confId string
	roster *roster.Roster

	handler handler.Handler
}

func NewBasicConference(confId string, localNode net.Node) (*BasicConference) {
	pbPack := &protocol.ProtobufPack{}

	conf := &BasicConference{confId, roster.NewRoster(),nil}

	msgHandler := handler.NewMessageHandler(conf,localNode,pbPack)
	conf.handler = msgHandler

	return conf
}

func (conf *BasicConference)ConferenceId() (string) {
	return conf.confId
}

func (conf *BasicConference) Roster() (*roster.Roster) {
	return conf.roster
}

func (conf *BasicConference) JoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest) ccs.CCSErrorCode {
	participant := conf.roster.GetParticipant(message.UserId)
	if participant != nil {
		conf.handler.OnAfterJoinConference(message.UserId, id, ccs.CCSErrorCode_AlreadyInMeeting)
		return ccs.CCSErrorCode_Success
	}

	participant = roster.NewParticipant(id, message.UserId, message.Name)
	conf.roster.AddParticipant(participant)

	conf.handler.OnAfterJoinConference(message.UserId, id, ccs.CCSErrorCode_Success)
	return ccs.CCSErrorCode_Success
}

func (conf *BasicConference) LeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest) ccs.CCSErrorCode{
	participant := conf.roster.GetParticipant(message.UserId)
	if participant == nil {
		conf.handler.OnAfterLeaveConference(nil, message.UserId, id, ccs.CCSErrorCode_NotInMeeting)
		return ccs.CCSErrorCode_Success
	}

	conf.roster.RemoveParticipant(participant)
	conf.handler.OnAfterLeaveConference(participant, message.UserId, id, ccs.CCSErrorCode_Success)
	return ccs.CCSErrorCode_Success
}

func (conf *BasicConference) EndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest) ccs.CCSErrorCode {
	conf.handler.OnAfterEndConference(message.UserId, id, ccs.CCSErrorCode_Success)
	return ccs.CCSErrorCode_Success
}

func (conf *BasicConference) Chat(id peer.ID, message *dvc_protocol.ChatMessage) {
	participant := conf.roster.GetParticipant(message.ToUserId)
	conf.handler.OnAfterChat(participant, message)
}