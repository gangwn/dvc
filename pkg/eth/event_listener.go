package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventCallback func(types.Log)

type EventListener interface {
	SetEventInfo(abiContent string, eventName string) (error)
	Subscribe(contactAddr common.Address, cb EventCallback) (error)
	Unsubscribe()
}