package config

import (
	"encoding/json"
	"os"
)

// Config contains configuration
type Config struct {
	Net net `json:"net"`
	Key KeyJson `json:"key""`
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
