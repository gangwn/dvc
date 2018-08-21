package net

import (
		"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-host"

	"github.com/gangwn/dvc/pkg/protocol"
			)

// MessageHandler processes the DVC message received from P2P network
type MessageHandler = func(peer.ID, *dvc_protocol.DVCMessage)

// Node is an object in the DVC P2P network
type Node interface {
	// ID returns local peer id for this node
	ID() peer.ID

	// Host returns libp2p host
	Host() host.Host

	// SetMessageHandler set message handler
	SetMessageHandler(msgHandler MessageHandler)

	// Connect connects the address and returns the peer id
	Connect(addr string) (peer.ID, error)

	// SendMessage sends message to a target peer id
	SendMessage(id peer.ID, message *dvc_protocol.DVCMessage) error
}