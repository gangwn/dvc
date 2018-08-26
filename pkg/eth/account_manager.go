package eth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type AccountManager interface {
	Account() accounts.Account

	KeyStore() *keystore.KeyStore

	Unlock(password string) error
}