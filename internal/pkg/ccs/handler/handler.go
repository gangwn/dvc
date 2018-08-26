package handler

import (
	"github.com/libp2p/go-libp2p-peer"
	"github.com/gangwn/dvc/internal/pkg/ccs"
	"github.com/gangwn/dvc/internal/pkg/ccs/roster"
)

type Handler interface {
	OnAfterJoinConference(userId string, id peer.ID, result ccs.CCSErrorCode)
	OnAfterLeaveConference(participant *roster.Participant, userId string, id peer.ID, result ccs.CCSErrorCode)
	OnAfterEndConference(userId string, id peer.ID, result ccs.CCSErrorCode)
	OnAfterChat(id peer.ID, result ccs.CCSErrorCode)
}