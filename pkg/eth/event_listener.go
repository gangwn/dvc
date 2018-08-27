package eth

import (
	"github.com/ethereum/go-ethereum/common"
		)

type EventListener interface {
	EventId(abiContent string, eventName string) (common.Hash, error)
	Subscribe(eventName string, contactAddr common.Address, eventId common.Hash, retry bool)
	Listen()
}