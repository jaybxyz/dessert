package client

import (
	"github.com/cosmos/cosmos-sdk/codec"

	cdc "github.com/kogisin/dessert/codec"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

// ConfigWrapper implements a wrapper around both a Tendermint RPC client and a
// Cosmos SDK REST client that allows for essential data queries.
type ConfigWrapper struct {
	rpcClient   rpcclient.Client // Tendermint RPC node
	lcdEndpoint string           // Full node
	cdc         *codec.Codec
}

// New creates a new ConfigWrapper
func New(rpcEndpoint, lcdEndpoint string) (ConfigWrapper, error) {
	rpcClient := rpcclient.NewHTTP(rpcEndpoint, "/websocket")

	if err := rpcClient.Start(); err != nil {
		return ConfigWrapper{}, err
	}

	return ConfigWrapper{rpcClient, lcdEndpoint, cdc.Codec}, nil
}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (cp ConfigWrapper) LatestHeight() (int64, error) {
	status, err := cp.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}
