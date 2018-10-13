package eth

import (
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type CCS struct {
	Addr common.Address
	IP string
	Port *big.Int
	PeerId string
}

type Client interface {

	EthClient() *ethclient.Client

	SetUp(password string, gasLimit uint64, gasPrice *big.Int) error

	ConferenceServiceContractAddress() common.Address

	CCSServiceContractAddress() common.Address

	RegisterCCS(ip string, port *big.Int, peerId string) error

	ListAllCCS() (error, []CCS)

	GetCCS(confId string) (error, CCS)

	NewJob(confId string, ccsAddr common.Address) error

	ScheduleConference(confId string, topic string, startTime *big.Int, duration *big.Int, invitees []common.Address) error

	CompleteJob() error
}