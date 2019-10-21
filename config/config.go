package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Config defines all necessary configuration parameters.
type Config struct {
	Client ClientConfig `toml:"client"`
	Wallet WalletConfig `toml:"wallet"`
}

// ClientConfig defines all client configuration parameters.
type ClientConfig struct {
	RPCEndpoint string `toml:"rpc_endpoint"`
	LCDEndpoint string `toml:"lcd_endpoint"`
}

// WalletConfig defines all wallet configuration parameters.
type WalletConfig struct {
	Network  string `toml:"network"`
	Mnemonic string `toml:"mnemonic"`
}

// ParseConfig attempts to read and parse a Juno config from the given file path.
// An error reading or parsing the config results in a panic.
func ParseConfig(configPath string) Config {
	if configPath == "" {
		log.Fatal("invalid configuration file")
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	var cfg Config
	if _, err := toml.Decode(string(configData), &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode config"))
	}

	return cfg
}
