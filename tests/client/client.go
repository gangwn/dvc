package main

import (
	"github.com/gangwn/dvc/pkg/util"
	"github.com/gangwn/dvc/pkg/net/basic"
	"github.com/gangwn/dvc/internal/pkg/ccs/config"
	"github.com/gangwn/dvc/pkg/protocol"
	"github.com/satori/go.uuid"
)

func main() {
	lisAddr :=[] config.ListenAddress{{"127.0.0.1", 8101}}
	addr := util.NewMultiaddr(lisAddr)

	_, priv, err := util.LoadKeyPair("/Users/gangwan/client")
	if err != nil {
		return
	}

	node := basic.NewBasicNode(addr, priv)
	if node == nil {
		return
	}

	id, err := node.Connect("/ip4/127.0.0.1/tcp/8100/ipfs/QmVzDgD8kiU57EPSZBYV45bzHYEiRXhQMats1ENHMfELVR")

	message := &dvc_protocol.DVCMessage{}
	message.Type = dvc_protocol.DVCMessage_JoinConference
	message.JoinConfMsg = &dvc_protocol.JoinConferenceRequest{}

	message.JoinConfMsg.UserId = uuid.NewV4().String()
	message.JoinConfMsg.ConferenceId = uuid.NewV4().String()
	message.JoinConfMsg.Name = "client1"

	node.SendMessage(id, message)

	select {}
}