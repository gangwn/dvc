package basiceth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"github.com/golang/glog"
	"github.com/gangwn/dvc/pkg/eth"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/cenkalti/backoff"
	"time"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

type BasicEventListener struct {
	ethClient eth.Client
	events     map[string]*Event
}

type Event struct {
	sub       ethereum.Subscription
	eventCh    chan types.Log
}

func NewBasicEventListener() *BasicEventListener {
	return &BasicEventListener{}
}

func (el *BasicEventListener)SubscribeMeetingScheduled() {

}

func (el *BasicEventListener)EventId(abiContent string, eventName string) (common.Hash, error) {
	abiJSON, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		return common.Hash{}, err
	}

	eventId := abiJSON.Events[eventName].Id()
	return eventId, nil
}

func (el *BasicEventListener) Subscribe(eventName string, contactAddr common.Address, eventId common.Hash, retry bool) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contactAddr},
		Topics:    [][]common.Hash{[]common.Hash{eventId}},
	}

	eventCh := make(chan types.Log)

	subscribe := func() error {
		sub, err := el.ethClient.EthClient().SubscribeFilterLogs(context.Background(), query, eventCh)
		if err != nil {
			glog.Errorf("Subscribe failed: %v", err)
			return err
		} else {
			glog.Infof("Subscribe successfully")
		}

		el.events[eventName] = &Event{
			sub:    sub,
			eventCh: eventCh,
		}

		return nil
	}

	if err := backoff.Retry(subscribe, backoff.NewConstantBackOff(time.Second * 2)); err != nil {
		glog.Errorf("Subscribe failed: %v", err)
		return err
	}

	return nil
}