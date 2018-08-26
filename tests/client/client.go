package main

import (
	"github.com/gangwn/dvc/pkg/util"
	"github.com/gangwn/dvc/pkg/net/basic"
	"github.com/gangwn/dvc/internal/pkg/ccs/config"
	"github.com/satori/go.uuid"
	"github.com/gangwn/dvc/pkg/protocol"
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

	pbPack := &protocol.ProtobufPack{}
	message := pbPack.CreateJoinConferenceRequest(uuid.NewV4().String(), uuid.NewV4().String(),  "client1")
	node.SendMessage(id, message)

	select {}
}