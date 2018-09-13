package events

import (
	"github.com/gangwn/dvc/pkg/eth"
	"github.com/gangwn/dvc/pkg/eth/basic"
	"github.com/gangwn/dvc/pkg/eth/contracts"
	"github.com/golang/glog"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
		)

type ConferenceScheduledEventHandler = func(eventData *ConferenceScheduledEventData)

type ConferenceScheduledEvent struct {
	eventListener eth.EventListener
	ethClient eth.Client
	eventData *ConferenceScheduledEventData
	Handler ConferenceScheduledEventHandler
}

type ConferenceScheduledEventData struct {
	CreatorAddr common.Address
	ConfId      string
	Topic       string
	StartTime   *big.Int
	Duration    *big.Int
	Invitees    []common.Address
}


func NewConferenceScheduledEvent(ethClient eth.Client) *ConferenceScheduledEvent {
	cse := &ConferenceScheduledEvent{
		basiceth.NewBasicEventListener(ethClient),ethClient, nil,nil,
	}

	err := cse.eventListener.SetEventInfo(contracts.ConferenceServiceABI, "ConferenceScheduled")
	if err != nil {
		glog.Errorf("Error SetEventInfo for event ConferenceScheduled")
		return nil
	}

	return cse
}



func (cse *ConferenceScheduledEvent) Listen() error {
	err := cse.eventListener.Subscribe(cse.ethClient.ConferenceServiceContractAddress(), func(log types.Log){
		abiJSON, _ := abi.JSON(strings.NewReader(contracts.ConferenceServiceABI))

		var ventData ConferenceScheduledEventData
		err := abiJSON.Unpack(&ventData, "ConferenceScheduled", log.Data)
		if err != nil {
			glog.Errorf("Failed to unpack ConferenceScheduled event data, error: %v", err)
			return
		}

		cse.eventData = &ventData
		glog.Infof("Unpack ConferenceScheduled event data, CreatorAddr: %v, ConfId: %v, Topic: %v, " +
			"StartTime: %v, Duration: %v, Invitees: %v", cse.eventData.CreatorAddr.Hex(), cse.eventData.ConfId,  cse.eventData.Topic,
			*cse.eventData.StartTime, *cse.eventData.Duration, cse.eventData.Invitees)
		cse.Handler(cse.eventData)
	})

	return err
}