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
	"fmt"
)


var EmptyEventId = fmt.Errorf("Empty EventId")

type BasicEventListener struct {
	ethClient eth.Client
	event     *Event
}

type Event struct {
	eventName string
	eventId common.Hash
	sub     ethereum.Subscription
	eventCh    chan types.Log
}

func NewBasicEventListener(ethClient eth.Client) *BasicEventListener {
	return &BasicEventListener{ethClient, nil}
}

func (el *BasicEventListener)SetEventInfo(abiContent string, eventName string) (error) {
	eid, err := el.eventId(abiContent, eventName)
	if err != nil {
		return err
	}

	el.event = &Event {
		eventName:    eventName,
		eventId: eid,
	}

	return nil
}

func (el *BasicEventListener) Subscribe(contactAddr common.Address, cb eth.EventCallback) (error) {
	if (el.event.eventId == common.Hash{}) {
		return EmptyEventId
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contactAddr},
		Topics:    [][]common.Hash{[]common.Hash{el.event.eventId}},
	}

	eventCh := make(chan types.Log)

	subscribe := func() error {
		//ctx, cancel := context.WithCancel(context.Background())
		//defer cancel()
		ctx := context.Background()
		sub, err := el.ethClient.EthClient().SubscribeFilterLogs(ctx, query, eventCh)
		if err != nil {
			glog.Errorf("Subscribe %v failed: %v", el.event.eventName, err)
			return err
		} else {
			glog.Infof("Subscribe %v successfully", el.event.eventName)
		}

		el.event.sub = sub
		el.event.eventCh = eventCh

		return nil
	}

	if err := backoff.Retry(subscribe, backoff.NewConstantBackOff(time.Second * 2)); err != nil {
		glog.Errorf("Subscribe %v failed: %v", el.event.eventName, err)
		return err
	}

	go el.Listen(cb, func() {
		glog.Infof("Retry to subscribe %v", el.event.eventName)

		if err := backoff.Retry(subscribe, backoff.NewConstantBackOff(time.Second*2)); err != nil {
			glog.Errorf("Error Retry to subscribe %v: %v", el.event.eventName, err)
			return
		}
	})

	return nil
}

func (el *BasicEventListener) Unsubscribe() {
	if el.event.sub == nil {
		return
	}

	close(el.event.eventCh)
	el.event.sub.Unsubscribe()

	el.event.eventCh = nil
	el.event.sub = nil
}

func (el *BasicEventListener)eventId(abiContent string, eventName string) (common.Hash, error) {
	abiJSON, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		return common.Hash{}, err
	}

	eventId := abiJSON.Events[eventName].Id()
	return eventId, nil
}

type ErrorCallback func()
func (el *BasicEventListener) Listen(cb eth.EventCallback, errCb ErrorCallback) {
	for {
		glog.Info("Start select event")
		select {
		case log, ok := <- el.event.eventCh:
			if !ok {
				glog.Errorf("Logs channel has been closed")
				return
			}

			glog.Info("Receive event logs")

			cb(log)
		case err := <-el.event.sub.Err():
			glog.Errorf("Log subscription error: %v", err)

			errCb()
		}
	}
}