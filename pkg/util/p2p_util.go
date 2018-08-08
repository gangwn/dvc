package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"peermeet/pkg/config"

	"github.com/golang/glog"
	crypto "github.com/libp2p/go-libp2p-crypto"
	ma "github.com/multiformats/go-multiaddr"
)

// NewMultiaddr convert IP and port to Multiaddr
func NewMultiaddr(addrs []config.ListenAddress) []ma.Multiaddr {
	var maddrs []ma.Multiaddr
	maddrs = make([]ma.Multiaddr, 0)

	for _, addr := range addrs {
		m := fmt.Sprintf("/ip4/%v/tcp/%d", addr.IP, addr.Port)
		maddr, err := ma.NewMultiaddr(m)
		if err != nil {
			glog.Errorf("Error NewMultiaddr %v to multiaddr: %v", m, err)
			continue
		}
		maddrs = append(maddrs, maddr)
	}

	return maddrs
}

func LoadKeyPair(keyDir string) (crypto.PubKey, crypto.PrivKey, error) {
	var pubKey crypto.PubKey
	var privKey crypto.PrivKey
	pubKey, privKey, err := loadKeyPairFromFile(keyDir)
	if err != nil || pubKey == nil || privKey == nil {
		privKey, pubKey, err = crypto.GenerateKeyPair(crypto.RSA, 2048)
		if err == nil {
			saveKeyPair(pubKey, privKey, keyDir)
		} else {
			glog.Errorf("Fail to generate key pair: %v", err)
			err = errors.New("GenerateKeyPairError")
		}
	}

	return pubKey, privKey, err
}

type KeyFile struct {
	PubKey  string
	PrivKey string
}

func loadKeyPairFromFile(keyDir string) (crypto.PubKey, crypto.PrivKey, error) {
	if keyDir == "" {
		return nil, nil, errors.New("EmptyKeyDir")
	}

	file, err := ioutil.ReadFile(path.Join(keyDir, "keys.json"))
	if err != nil {
		return nil, nil, err
	}

	var keyFile KeyFile
	err = json.Unmarshal(file, &keyFile)
	if err != nil {
		return nil, nil, err
	}

	pubBytes, err := crypto.ConfigDecodeKey(keyFile.PubKey)
	if err != nil {
		return nil, nil, err
	}

	privBytes, err := crypto.ConfigDecodeKey(keyFile.PrivKey)
	if err != nil {
		return nil, nil, err
	}

	pubKey, err := crypto.UnmarshalPublicKey(pubBytes)
	if err != nil {
		return nil, nil, err
	}

	privKey, err := crypto.UnmarshalPrivateKey(privBytes)
	if err != nil {
		return nil, nil, err
	}

	return pubKey, privKey, nil
}

func saveKeyPair(pubKey crypto.PubKey, privKey crypto.PrivKey, keyDir string) {
	if keyDir == "" {
		glog.Errorf("Can't save keypair, keyDir is empty!")
		return
	}

	pubBytes, _ := pubKey.Bytes()
	privBytes, _ := privKey.Bytes()
	keyFile := KeyFile{crypto.ConfigEncodeKey(pubBytes), crypto.ConfigEncodeKey(privBytes)}
	keyFileJson, err := json.Marshal(keyFile)
	if err != nil {
		glog.Errorf("Fail to serialize the key store to json format: %v", err)
		return
	}

	keyPath := path.Join(keyDir, "keys.json")
	err = ioutil.WriteFile(keyPath, keyFileJson, 0644)
	if err != nil {
		glog.Errorf("Fail to save the key to file %v, error: %v", keyPath, err)
		return
	}
}
