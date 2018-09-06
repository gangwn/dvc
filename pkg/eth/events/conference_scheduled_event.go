package events

import (
	"github.com/gangwn/dvc/pkg/eth"
	"github.com/gangwn/dvc/pkg/eth/basic"
	"github.com/gangwn/dvc/pkg/eth/contracts"
	"github.com/golang/glog"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	)

type ConferenceScheduledEvent struct {
	eventListener eth.EventListener
	ethClient eth.Client
	eventData *contracts.ConferenceServiceConferenceScheduled
}

func NewConferenceScheduledEvent(ethClient eth.Client) *ConferenceScheduledEvent {
	cse := &ConferenceScheduledEvent{
		basiceth.NewBasicEventListener(ethClient),ethClient, nil,
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
		err := abiJSON.Unpack(&cse.eventData, "ConferenceScheduled", log.Data)
		if err != nil {
			glog.Errorf("Failed to unpack ConferenceScheduled event data: %v, error: %v", log, err)
		}

		glog.Info("Unpack ConferenceScheduled event data: %v", *cse.eventData)
	})

	return err
}