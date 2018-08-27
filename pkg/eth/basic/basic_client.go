package basiceth

// Generate contract code via below commands
//go:generate abigen --sol  ../../../api/contracts/WebexMeeting.sol --pkg contracts --out ../contracts/webex_meeting.go

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gangwn/dvc/pkg/eth/contracts"
	"github.com/gangwn/dvc/pkg/eth"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"errors"
	"github.com/golang/glog"
)

type BasicClient struct {
	ethClient *ethclient.Client
	webexMeetingContractAddr common.Address

	webexMeetingContract *contracts.WebexMeeting
	webexMeetingSession *contracts.WebexMeetingSession

	am eth.AccountManager
}

func NewBasicClient(account string, keystoreDir string, gethRPCPath string, contactAddr string) (*BasicClient, error) {
	am, err := NewBasicAccountManager(account, keystoreDir)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(gethRPCPath)
	if err != nil {
		return nil, err
	}

	basicClient := &BasicClient{
		ethClient:client,
		webexMeetingContractAddr: common.HexToAddress(contactAddr),
		am: am}

	return basicClient, nil
}


func (client *BasicClient) EthClient() *ethclient.Client{
	return client.ethClient
}


func (client *BasicClient) SetUp(password string, gasLimit uint64, gasPrice *big.Int) error {
	err := client.am.Unlock(password)
	if err != nil {
		return err
	}

	webexMeetingContract, err := contracts.NewWebexMeeting(client.webexMeetingContractAddr, client.ethClient)
	if err != nil {
		return err
	}
	client.webexMeetingContract = webexMeetingContract

	transactOpts, err := client.createTransactOpts(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	client.webexMeetingSession = &contracts.WebexMeetingSession {
		client.webexMeetingContract,
		bind.CallOpts{
			Pending: true,
		},
		*transactOpts,
	}

	return nil
}

func (client *BasicClient) ScheduleConference() {
	glog.Infof("ethClient.ScheduleConference")

	//_, err := client.webexMeetingSession.ScheduleMeeting("top",  big.NewInt(11),  big.NewInt(111))
	//if err != nil {
	//	glog.Infof("ScheduleMeeting error: %v", err)
	//}
	//
	//key,err := client.webexMeetingSession.MMeetingKey()
	//glog.Infof("key: %v, err:%s", key, err)
	tx, err:= client.webexMeetingSession.StartMeeting()
	glog.Infof("tx: %v, err:%s", tx, err)
}


func (client *BasicClient) createTransactOpts(gasLimit uint64, gasPrice *big.Int) (*bind.TransactOpts, error) {
	return &bind.TransactOpts{
		From:     client.am.Account().Address,
		GasLimit: gasLimit,
		GasPrice: gasPrice,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != client.am.Account().Address {
				return nil, errors.New("not authorized to sign this account")
			}

			signature, err := client.am.KeyStore().SignHash(client.am.Account(), signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}

			return tx.WithSignature(signer, signature)
		},
	}, nil
}