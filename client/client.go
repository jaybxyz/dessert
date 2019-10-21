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
	Codec       *codec.Codec
}

// New creates a new ConfigWrapper
func New(rpcEndpoint, lcdEndpoint string) (ConfigWrapper, error) {
	rpcClient := rpcclient.NewHTTP(rpcEndpoint, "/websocket")

	if err := rpcClient.Start(); err != nil {
		return ConfigWrapper{}, err
	}

	return ConfigWrapper{rpcClient, lcdEndpoint, cdc.Codec}, nil
}

// Stop defers the node stop execution to the RPC client.
func (cw ConfigWrapper) Stop() error {
	return cw.rpcClient.Stop()
}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (cw ConfigWrapper) LatestHeight() (int64, error) {
	status, err := cw.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}
