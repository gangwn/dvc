package controler

import (
	"github.com/gangwn/dvc/pkg/net"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
	"github.com/gangwn/dvc/internal/pkg/ccs/conference"
	"github.com/golang/glog"
)

type Controler struct {
	node net.Node
	roster *roster.Roster
	conferences map[string]*conference.Conference
}


func NewControler(node net.Node) (*Controler) {
	controler := &Controler{node, nil, make(map[string]*conference.Conference)}
	controler.node.SetMessageHandler(controler.MessageHandler)
	return controler;
}

func (controler *Controler) MessageHandler(id peer.ID, message *dvc_protocol.DVCMessage) {
	glog.Infof("Receive message type %s from id %s", message.Type, id.Pretty() )


	if message.GetType() == dvc_protocol.DVCMessage_JoinConference {
		//controler.conference.JoinConference(message.GetJoinConfMsg())
	}
}

func (controler *Controler) handleJoinConference(message *dvc_protocol.JoinConferenceMessage) {

}