package basicconference

import (
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/net"
)

type BasicConference struct {
	confId string
	roster *roster.Roster
	localNode net.Node
}

func NewBasicConference(confId string, localNode net.Node) (conf *BasicConference) {
	return &BasicConference{confId, roster.NewRoster(), localNode}
}

func (conf *BasicConference) JoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest) {
	participant := conf.roster.GetParticipant(message.UserId)
	if participant != nil {
		conf.sendJoinConferenceResponse(id, 0)
		return
	}

	participant = roster.NewParticipant(id, message.UserId, message.Name)
	conf.roster.AddParticipant(participant)

	conf.sendJoinConferenceResponse(id, 0)

	conf.NotifyParticipantJoin(participant)
}

func (conf *BasicConference) LeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest) {

}

func (conf *BasicConference) EndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest) {

}

func (conf *BasicConference) Chat(id peer.ID, message *dvc_protocol.ChatMessage) {

}

func (conf *BasicConference) sendJoinConferenceResponse(id peer.ID, result int32) {
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_JoinConferenceResponse
	message.JoinConfRsp = &dvc_protocol.JoinConferenceResponse{}
	message.JoinConfRsp.Result = result

	conf.localNode.SendMessage(id, message)
}


func (conf *BasicConference) NotifyParticipantJoin( participant *roster.Participant) {
	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_RosterMessage
	message.RosterMsg = &dvc_protocol.RosterMessage{}

	p := dvc_protocol.Participant{}
	p.State = dvc_protocol.Participant_Joined
	p.Name = participant.Name
	p.UserId = participant.UserId

	message.RosterMsg.Participants = make([]*dvc_protocol.Participant, 1)
	message.RosterMsg.Participants = append(message.RosterMsg.Participants, &p)

	conf.broadcast(message)
}

func (conf *BasicConference) broadcast(message* dvc_protocol.DVCMessage){
	participants := conf.roster.Participants()
	for userId := range participants {
		conf.localNode.SendMessage(participants[userId].Pid, message)
	}
}