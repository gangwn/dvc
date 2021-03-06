package basiceth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gangwn/dvc/pkg/eth"
	"github.com/gangwn/dvc/pkg/eth/contracts"
	"github.com/golang/glog"
	"math/big"
)

var EmptyEthDialAddress = errors.New("EmptyEthDialAddress")

type BasicClient struct {
	ethClient *ethclient.Client

	serviceManagerAddr     common.Address
	serviceManagerContract *contracts.ServiceManager
	serviceManagerSession  *contracts.ServiceManagerSession

	ccsServiceAddr     common.Address
	ccsServiceContract *contracts.CCSService
	ccsServiceSession  *contracts.CCSServiceSession

	conferenceServiceAddr     common.Address
	conferenceServiceContract *contracts.ConferenceService
	conferenceServiceSession  *contracts.ConferenceServiceSession

	fixedSupplyTokenAddr     common.Address
	fixedSupplyTokenContract *contracts.FixedSupplyToken
	fixedSupplyTokenSession  *contracts.FixedSupplyTokenSession

	am eth.AccountManager
}

func NewBasicClient(account string, keystoreDir string, gethRPCPath string, rpcUrl string, contactAddr string) (*BasicClient, error) {
	am, err := NewBasicAccountManager(account, keystoreDir)
	if err != nil {
		return nil, err
	}

	var client *ethclient.Client
	if rpcUrl != "" {
		client, err = ethclient.Dial(rpcUrl)
	} else if gethRPCPath != "" {
		client, err = ethclient.Dial(gethRPCPath)
	} else {
		err = EmptyEthDialAddress
	}

	if err != nil {
		return nil, err
	}

	basicClient := &BasicClient{
		ethClient:          client,
		serviceManagerAddr: common.HexToAddress(contactAddr),
		am:                 am}

	return basicClient, nil
}

func (client *BasicClient) EthClient() *ethclient.Client {
	return client.ethClient
}

func (client *BasicClient) SetUp(password string, gasLimit uint64, gasPrice *big.Int) error {
	err := client.am.Unlock(password)
	if err != nil {
		return err
	}

	err = client.setUpServiceManager(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	err = client.setUpCCSService(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	err = client.setUpConferenceService(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	err = client.setUpFixedSupplyToken(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	return nil
}

func (client *BasicClient) ConferenceServiceContractAddress() common.Address {
	return client.conferenceServiceAddr
}

func (client *BasicClient) CCSServiceContractAddress() common.Address {
	return client.ccsServiceAddr
}

func (client *BasicClient) RegisterCCS(ip string, port *big.Int, peerId string) error {
	_, err := client.ccsServiceSession.RegisterCCS(ip, port, peerId)
	glog.Infof("RegisterCCS, ip: %v, port: %v, peerId: %v, err: %v", ip, port, peerId, err)

	return err
}

func (client *BasicClient) ListAllCCS() (error, []eth.CCS) {
	glog.Infof("listAllCCS")
	addr, ip, port, peerId, err := client.ccsServiceSession.GetFirstCCS()
	if err != nil {
		return err, nil
	}

	var ccss []eth.CCS
	ccs := eth.CCS{addr, ip, port, peerId}
	ccss = append(ccss, ccs)

	nextAddr, ip, port, peerId, err := client.ccsServiceSession.GetNextCCS(addr)
	for (err == nil && nextAddr != addr && nextAddr != common.Address{0}) {
		ccs := eth.CCS{nextAddr, ip, port, peerId}
		ccss = append(ccss, ccs)

		nextAddr, ip, port, peerId, err = client.ccsServiceSession.GetNextCCS(nextAddr)
	}

	return err, ccss
}

func (client *BasicClient) GetCCS(confId string) (error, eth.CCS) {
	addr, ip, port, peerId, err := client.ccsServiceSession.GetCCS(confId)
	ccs := eth.CCS{addr, ip, port, peerId}

	glog.Infof("GetCCS, confId: %v, ccsAddr: %v", confId, ccs.Addr.Hex())

	return err, ccs
}

func (client *BasicClient) NewJob(confId string, ccsAddr common.Address) error {

	_, err := client.ccsServiceSession.NewJob(confId, ccsAddr)
	return err
}

func (client *BasicClient) ScheduleConference(confId string, topic string, startTime *big.Int, duration *big.Int,
	invitees []common.Address) error {

	glog.Infof("ScheduleConference, confId: %v, topic: %v, startTime: %v, duration: %v", confId, topic, *startTime, *duration)

	_, err := client.conferenceServiceSession.ScheduleConference(confId, topic, startTime, duration, invitees)

	if err != nil {
		glog.Infof("error ScheduleConference, error: %v", err)
	}

	return err
}

func (client *BasicClient) CompleteJob() error {
	_, err := client.fixedSupplyTokenSession.Transfer(client.am.Account().Address, big.NewInt(int64((1000))))
	return err
}

func (client *BasicClient) setUpServiceManager(gasLimit uint64, gasPrice *big.Int) error {
	glog.Infof("setUpServiceManager, address: %v", client.serviceManagerAddr.Hex())

	serviceManagerContract, err := contracts.NewServiceManager(client.serviceManagerAddr, client.ethClient)
	if err != nil {
		return err
	}
	client.serviceManagerContract = serviceManagerContract

	transactOpts, err := client.createTransactOpts(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	client.serviceManagerSession = &contracts.ServiceManagerSession{
		client.serviceManagerContract,
		bind.CallOpts{
			Pending: true,
		},
		*transactOpts,
	}

	return nil
}

func (client *BasicClient) setUpCCSService(gasLimit uint64, gasPrice *big.Int) error {
	ccsServiceAddr, err := client.serviceManagerSession.GetService("CCSService")
	if err != nil {
		return err
	}

	client.ccsServiceAddr = ccsServiceAddr

	glog.Infof("setUpCCSService, address: %v", ccsServiceAddr.Hex())

	ccsServiceContract, err := contracts.NewCCSService(ccsServiceAddr, client.ethClient)
	if err != nil {
		return err
	}
	client.ccsServiceContract = ccsServiceContract

	transactOpts, err := client.createTransactOpts(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	client.ccsServiceSession = &contracts.CCSServiceSession{
		client.ccsServiceContract,
		bind.CallOpts{
			Pending: true,
		},
		*transactOpts,
	}

	return nil
}

func (client *BasicClient) setUpFixedSupplyToken(gasLimit uint64, gasPrice *big.Int) error {
	fixedSupplyTokenAddr, err := client.serviceManagerSession.GetService("FixedSupplyToken")
	if err != nil {
		return err
	}

	client.fixedSupplyTokenAddr = fixedSupplyTokenAddr

	glog.Infof("setUpFixedSupplyToken, address: %v", fixedSupplyTokenAddr.Hex())

	fixedSupplyTokenContract, err := contracts.NewFixedSupplyToken(fixedSupplyTokenAddr, client.ethClient)
	if err != nil {
		return err
	}
	client.fixedSupplyTokenContract = fixedSupplyTokenContract

	transactOpts, err := client.createTransactOpts(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	client.fixedSupplyTokenSession = &contracts.FixedSupplyTokenSession{
		client.fixedSupplyTokenContract,
		bind.CallOpts{
			Pending: true,
		},
		*transactOpts,
	}

	return nil
}

func (client *BasicClient) setUpConferenceService(gasLimit uint64, gasPrice *big.Int) error {
	conferenceServiceAddr, err := client.serviceManagerSession.GetService("ConferenceService")
	if err != nil {
		return err
	}

	client.conferenceServiceAddr = conferenceServiceAddr

	glog.Infof("setUpConferenceService, address: %v", conferenceServiceAddr.Hex())

	conferenceServiceContract, err := contracts.NewConferenceService(conferenceServiceAddr, client.ethClient)
	if err != nil {
		return err
	}
	client.conferenceServiceContract = conferenceServiceContract

	transactOpts, err := client.createTransactOpts(gasLimit, gasPrice)
	if err != nil {
		return err
	}

	client.conferenceServiceSession = &contracts.ConferenceServiceSession{
		client.conferenceServiceContract,
		bind.CallOpts{
			Pending: true,
		},
		*transactOpts,
	}

	return nil
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
