package handler

import (
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/net"
	"github.com/gangwn/dvc/internal/pkg/ccs"
	"github.com/gangwn/dvc/internal/pkg/ccs/conference"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
)

type MessageHandler struct {
	conf      conference.Conference
	localNode net.Node

	pbPack *protocol.ProtobufPack

	Next Handler
}

func NewMessageHandler(conf conference.Conference, localNode net.Node, pbPack *protocol.ProtobufPack) (*MessageHandler) {
	msgHandler := &MessageHandler{conf,localNode,pbPack,nil}
	return msgHandler
}


func (handler *MessageHandler) OnAfterJoinConference(userId string, id peer.ID, result ccs.CCSErrorCode) {
	if result ==  ccs.CCSErrorCode_AlreadyInMeeting {
		handler.localNode.SendMessage(id, handler.pbPack.CreateJoinConferenceResponse(handler.conf.ConferenceId(),userId, 0))
	} else {
		handler.localNode.SendMessage(id, handler.pbPack.CreateJoinConferenceResponse(handler.conf.ConferenceId(),userId, int32(result)))
	}

	if result == ccs.CCSErrorCode_Success {
		participant := handler.conf.Roster().GetParticipant(userId)
		if participant != nil {

		}
	}

	if handler.Next != nil {
		handler.Next.OnAfterJoinConference(userId, id, result)
	}
}

func (handler *MessageHandler) OnAfterLeaveConference(participant *roster.Participant, userId string, id peer.ID, result ccs.CCSErrorCode) {
	handler.localNode.SendMessage(id, handler.pbPack.CreateLeaveConferenceResponse(handler.conf.ConferenceId(), userId, int32(result)))

	if result == ccs.CCSErrorCode_Success {
		handler.broadcast(handler.pbPack.CreateRosterMessage(handler.conf.ConferenceId(),participant,true))
	}

	if handler.Next != nil {
		handler.Next.OnAfterLeaveConference(participant, userId, id, result)
	}
}

func (handler *MessageHandler) OnAfterEndConference(userId string, id peer.ID, result ccs.CCSErrorCode) {
	handler.localNode.SendMessage(id, handler.pbPack.CreateEndConferenceResponse(handler.conf.ConferenceId(),userId, int32(result)))

	if handler.Next != nil {
		handler.Next.OnAfterEndConference(userId, id, result)
	}
}

func (handler *MessageHandler) OnAfterChat(to *roster.Participant, message* dvc_protocol.ChatMessage) {
	if to == nil {
		return
	}

	handler.localNode.SendMessage(to.Pid, handler.pbPack.CreateChatMessage(message))

	if handler.Next != nil {
		handler.Next.OnAfterChat(to, message)
	}
}

func (handler *MessageHandler) broadcast(message* dvc_protocol.DVCMessage){
	participants := handler.conf.Roster().Participants()
	for userId := range participants {
		handler.localNode.SendMessage(participants[userId].Pid, message)
	}
}