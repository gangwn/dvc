package basiceth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/golang/glog"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/ethereum/go-ethereum/console"
)

var AccountNotFound = fmt.Errorf("Account not found")
var ReadPassphraseFail = fmt.Errorf("Read passphrase fail")

type BasicAccountManager struct {
	account accounts.Account
	ks *keystore.KeyStore
}

func NewBasicAccountManager(account string, keystoreDir string) (*BasicAccountManager, error) {
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	am := &BasicAccountManager{
		ks:ks,
	}

	acc, err := am.tryGetAccount(account)
	if err != nil {
		return nil, err
	}

	glog.Infof("ETH account: %v", acc.Address.Hex())

	am.account = acc

	return am, nil
}

func (accountManager *BasicAccountManager) Account() accounts.Account {
	return accountManager.account
}

func (accountManager *BasicAccountManager) KeyStore() *keystore.KeyStore {
	return accountManager.ks
}

func (accountManager *BasicAccountManager) Unlock(password string) error {
	err := accountManager.ks.Unlock(accountManager.account, password)
	if err != nil {
		password, err = console.Stdin.PromptPassword("Enter passphrase to unlock ETH account: ")
		if err != nil {
			err = accountManager.ks.Unlock(accountManager.account, password)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (accountManager *BasicAccountManager) tryGetAccount(account string) (accounts.Account, error) {
	acc := accounts.Account{}
	var err error

	accountAddr := common.HexToAddress(account)
	hasAccount := accountManager.ks.HasAddress(accountAddr)
	if !hasAccount && account != "" {
		err = AccountNotFound
	} else if account == "" && len(accountManager.ks.Accounts()) > 0 {
		fmt.Printf("No avaiable account founded, try to create an ETH account")
		acc, err = accountManager.createAccount()
	} else if hasAccount || (account == "" &&  len(accountManager.ks.Accounts()) == 0) {
		acc, err = accountManager.getAccount(accountAddr)
	} else {
		err = AccountNotFound
	}

	return acc, err
}


func (accountManager *BasicAccountManager)  createAccount() (accounts.Account, error) {
	passphrase := getPassphrase()
	if passphrase == "" {
		return accounts.Account{}, ReadPassphraseFail
	}

	return accountManager.ks.NewAccount(passphrase)
}

func (accountManager *BasicAccountManager)  getAccount(accountAddr common.Address)(accounts.Account, error) {
	accs := accountManager.ks.Accounts()
	if (accountAddr == common.Address{}) && len(accs) > 0 {
		return accs[0], nil
	} else if (accountAddr != common.Address{}) {
		for _, account := range accs {
			if account.Address == accountAddr {
				return account, nil
			}
		}
	}
	return accounts.Account{}, AccountNotFound
}

func getPassphrase() (string) {
	passphrase, err := console.Stdin.PromptPassword("Enter passphrase: ")
	if err != nil {
		fmt.Printf("Passphrase can't be empty")
		return ""
	}

	passphraseConfirm, err := console.Stdin.PromptPassword("Repeat passphrase: ")
	if err != nil {
		fmt.Printf("Passphrase can't be empty")
		return ""
	}

	if passphrase != passphraseConfirm {
		fmt.Printf("Passphrase does not match the confirm passphrase")
		return ""
	}

	return passphrase
}