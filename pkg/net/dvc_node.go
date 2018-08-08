package net

import (
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	peer "github.com/libp2p/go-libp2p-peer"
	ma "github.com/multiformats/go-multiaddr"
)

// DVCNode represents a node in P2P network
type DVCNode struct {
	ID   peer.ID
	Host host.Host
}

// New constructs and sets up a new *DVCNode.
func NewDVCNode(listenAddrs []ma.Multiaddr, pubKey crypto.PubKey, privKey crypto.PrivKey) *DVCNode {

}
