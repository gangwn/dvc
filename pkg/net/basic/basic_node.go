package basic

import (
	"context"

	"github.com/gangwn/dvc/pkg/net"

	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p"
	inet "github.com/libp2p/go-libp2p-net"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/golang/glog"
	"github.com/gangwn/dvc/pkg/protocol/pb"
	"errors"
		"fmt"
	"github.com/libp2p/go-libp2p-peerstore"
)

var StreamNotFound = errors.New("StreamNotFound")
var ErrorNewStream = errors.New("ErrorNewStream")

// BasicNode represents a node in P2P network
type BasicNode struct {
	pid        peer.ID
	host       host.Host
	msgHandler net.MessageHandler
	outStreams map[peer.ID]*net.OutStream
}

// NewBasicNode is function to set up a new *BasicNode.
func NewBasicNode(listenAddrs []ma.Multiaddr, privKey crypto.PrivKey) *BasicNode {
	opts := []libp2p.Option{
		libp2p.ListenAddrs(listenAddrs[0]),
		libp2p.Identity(privKey),
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil
	}

	streams := make(map[peer.ID]*net.OutStream)
	node := &BasicNode{basicHost.ID(), basicHost, nil, streams}

	basicHost.SetStreamHandler(dvcProtocol, node.onMessageReceived)

	return node
}

func (node *BasicNode) ID() peer.ID {
	return node.pid
}

func (node *BasicNode) Host() host.Host {
	return node.host
}

func (node *BasicNode) SetMessageHandler(msgHandler net.MessageHandler) {
	node.msgHandler = msgHandler
}

func (node *BasicNode) SendMessage(pid peer.ID, message *dvc_protocol.DVCMessage) error {
	stream := node.getOutStream(pid)
	if stream == nil {
		return StreamNotFound
	}

	err := stream.SendMessage(message)
	if(err != nil) {
		glog.Errorf("Send message fail, message: %v, err: v", message, err)
	}

	return err
}

func (node *BasicNode) Connect(addr string) (peer.ID, error) {
	glog.Infof("Connect to remote: %v", addr)

	pid, err := node.AddAddrToPeerstore(addr)
	if err != nil {
		return pid, err
	}

	stream := node.getOutStream(pid)
	if stream == nil {
		return pid, ErrorNewStream
	}

	return pid,nil
}

func (node *BasicNode) AddAddrToPeerstore(addr string) (peer.ID, error) {
	ipfsaddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return "", err
	}

	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		return "", err
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		return "", err
	}

	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	node.host.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)
	return peerid, nil
}

const dvcProtocol = "/dvc/0.0.1"

func (node *BasicNode)onMessageReceived(s inet.Stream) {
	inStream := net.NewInStream(s)
	for {
		message := inStream.DecodeMessage()
		if message == nil {
			break
		}

		glog.Infof("Receive message from %s to local peer %s, message type", s.Conn().RemotePeer().Pretty(),
			s.Conn().LocalPeer().Pretty(), message.GetType())

		if node.msgHandler != nil {
			node.msgHandler(s.Conn().RemotePeer(), message)
		}
	}

	glog.Infof("Close stream id: %v", s.Conn().RemotePeer().Pretty())

	inStream.Close()
}

func (node *BasicNode) addStream(s inet.Stream) bool {
	if _, ok := node.outStreams[s.Conn().RemotePeer()]; ok {
		return false
	}

	stream := net.NewOutStream(s)
	node.outStreams[s.Conn().RemotePeer()] = stream

	return true
}

func (node *BasicNode) getOutStream(id peer.ID) *net.OutStream {
	if stream, ok := node.outStreams[id]; ok {
		return stream
	}

	stream := node.newOutStream(id)
	if stream == nil {
		return nil
	}

	node.outStreams[id] = stream

	return stream
}

func (node *BasicNode) newOutStream(pid peer.ID) *net.OutStream {
	s, err := node.host.NewStream(context.Background(), pid, dvcProtocol)
	if err != nil {
		glog.Errorf("Error creating out stream, local: %v, remote %v: %v", node.ID().Pretty(), pid.Pretty(), err)
		return nil
	}

	outStream := net.NewOutStream(s)

	glog.Infof("Create out stream, peer id: %v", pid.Pretty())
	return outStream
}