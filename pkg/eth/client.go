package eth

import (
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {

	EthClient() *ethclient.Client

	SetUp(password string, gasLimit uint64, gasPrice *big.Int)

	RegisterCCS(ip string, port *big.Int)

	ScheduleConference()
}