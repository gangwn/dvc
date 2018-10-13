package controler

import (
	"github.com/gangwn/dvc/pkg/net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"github.com/golang/glog"
	"github.com/gangwn/dvc/internal/pkg/ccs/conference"
	"github.com/gangwn/dvc/internal/pkg/ccs/conference/basic"
	"github.com/gangwn/dvc/internal/pkg/ccs"
	"github.com/gangwn/dvc/pkg/eth"
)

type Controler struct {
	localNode net.Node
	conferences map[string]conference.Conference
	ethClient eth.Client
}

func NewControler(node net.Node) (*Controler) {
	controler := &Controler{node, make(map[string]conference.Conference), nil}
	controler.localNode.SetMessageHandler(controler.MessageHandler)
	return controler;
}

func(controler *Controler) SetEthClient(ethClient eth.Client) {
	controler.ethClient = ethClient;
}

func(controler *Controler) RemoveConference(conference conference.Conference) {
	if _, ok := controler.conferences[conference.ConferenceId()]; ok {
		delete(controler.conferences, conference.ConferenceId())
	}
}

func (controler *Controler) MessageHandler(id peer.ID, message *dvc_protocol.DVCMessage) {
	glog.Infof("Receive message type %s from id %s", message.Type, id.Pretty() )

	if message.GetType() == dvc_protocol.DVCMessage_JoinConferenceRequest {
		controler.handleJoinConference(id, message.GetJoinConfReq())
	} else if message.GetType() == dvc_protocol.DVCMessage_LeaveConferenceRequest {
		controler.handleLeaveConference(id, message.GetLeaveConfReq())
	} else if message.GetType() == dvc_protocol.DVCMessage_EndConferenceRequest {
		controler.handleEndConference(id, message.GetEndConfReq())
	} else if message.GetType() == dvc_protocol.DVCMessage_Chat {
		controler.handleChat(id, message.GetChatMsg())
	}
}

func (controler *Controler) handleJoinConference(id peer.ID, message *dvc_protocol.JoinConferenceRequest) {
	if conf, ok := controler.conferences[message.ConferenceId]; ok {
		conf.JoinConference(id, message)
	} else {
		conf := basicconference.NewBasicConference(message.ConferenceId, controler.localNode, controler.ethClient)
		conf.JoinConference(id, message)
		
		controler.conferences[message.ConferenceId] = conf
	}
}

func (controler *Controler) handleLeaveConference(id peer.ID, message *dvc_protocol.LeaveConferenceRequest) {
	if conf, ok := controler.conferences[message.ConferenceId]; ok {
		conf.LeaveConference(id, message)
	}
}

func (controler *Controler) handleEndConference(id peer.ID, message *dvc_protocol.EndConferenceRequest) {
	if conf, ok := controler.conferences[message.ConferenceId]; ok {
		result := conf.EndConference(id, message)
		if result == ccs.CCSErrorCode_Success {
			delete(controler.conferences, message.ConferenceId)
		}
	}
}

func (controler *Controler) handleChat(id peer.ID, message *dvc_protocol.ChatMessage) {
	if conf, ok := controler.conferences[message.ConferenceId]; ok {
		conf.Chat(id, message)
	}
}