package protocol

import (
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
)

type ProtobufPack struct {

}

func (*ProtobufPack) CreateJoinConferenceRequest(confId string, userId string, name string) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_JoinConferenceRequest
	message.JoinConfReq = &dvc_protocol.JoinConferenceRequest{}

	message.JoinConfReq.UserId = userId
	message.JoinConfReq.ConferenceId = confId
	message.JoinConfReq.Name = name

	return message
}

func (*ProtobufPack) CreateJoinConferenceResponse(confId string, userId string, result int32) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_JoinConferenceResponse
	message.JoinConfRsp = &dvc_protocol.JoinConferenceResponse{}

	message.JoinConfRsp.UserId = userId
	message.JoinConfRsp.Result = result
	message.JoinConfRsp.ConferenceId = confId

	return message
}


func (*ProtobufPack) CreateLeaveConferenceResponse(confId string, userId string, result int32) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_LeaveConferenceResponse
	message.LeaveConfRsp = &dvc_protocol.LeaveConferenceResponse{}

	message.LeaveConfRsp.UserId = userId
	message.LeaveConfRsp.Result = result
	message.LeaveConfRsp.ConferenceId = confId

	return message
}


func (*ProtobufPack) CreateEndConferenceResponse(confId string, userId string, result int32) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_EndConferenceResponse
	message.EndConfRsp = &dvc_protocol.EndConferenceResponse{}

	message.EndConfRsp.UserId = userId
	message.EndConfRsp.Result = result
	message.EndConfRsp.ConferenceId = confId

	return message
}

func (*ProtobufPack) CreateRosterMessage(confId string, participant *roster.Participant, joined bool) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_RosterMessage
	message.RosterMsg = &dvc_protocol.RosterMessage{}

	p := dvc_protocol.Participant{}
	if joined {
		p.State = dvc_protocol.Participant_Joined
	} else {
		p.State = dvc_protocol.Participant_Left
	}

	p.Name = participant.Name
	p.UserId = participant.UserId

	message.RosterMsg.Participants = make([]*dvc_protocol.Participant, 1)
	message.RosterMsg.Participants = append(message.RosterMsg.Participants, &p)
	message.RosterMsg.ConferenceId = confId

	return message
}

func (*ProtobufPack) CreateChatMessage(chatMessage *dvc_protocol.ChatMessage) (*dvc_protocol.DVCMessage){
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_Chat
	message.ChatMsg = chatMessage

	return message
}
