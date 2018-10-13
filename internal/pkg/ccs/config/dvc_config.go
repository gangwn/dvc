package config

import (
	"encoding/json"
	"os"
)

// Config contains configuration
type Config struct {
	Net net `json:"net"`
	Key KeyJson `json:"key""`
	Eth eth `json:"eth"`
}

type net struct {
	LisAddrs []ListenAddress `json:"listen_addresses"`
}

type ListenAddress struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type KeyJson struct {
	Dir   string `json:"directory"`
}

type eth struct {
	GethRPCPath string `json:"gethRPCPath"`
	RPCAddress string `json:"rpcAddress"`
	RPCPort int `json:"rpcPort"`
	KeystoreDir string `json:"keystoreDir"`
	EthAccount string `json:"ethAccount"`
	ContractAddr string `json:"contractAddr"`
	GasLimit string `json:"gasLimit"`
	GasPrice int `json:"gasPrice"`
	Password string `json:"password"`
}

// ReadConfig read the configuration from file
func ReadConfig(cfgPath string) (*Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
