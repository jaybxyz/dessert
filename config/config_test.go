package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cfgFile = "../config.toml"
)

func TestParseConfig(t *testing.T) {
	cfg := ParseConfig(cfgFile)

	if assert.NotEmpty(t, cfg) {
		t.Logf("RPCEndpoint %s", cfg.Client.LCDEndpoint)
		t.Logf("LCDEndpoint %s", cfg.Client.LCDEndpoint)
		t.Logf("Mnemonic %s", cfg.Wallet.Mnemonic)
	}
}
