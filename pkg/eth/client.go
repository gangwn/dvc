package eth

import "math/big"

type Client interface {
	Initialize(password string, gasLimit uint64, gasPrice *big.Int)

	ScheduleConference()
}