package wallet

import (
	"github.com/rs/zerolog/log"

	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/bech32"
)

// GetBech32AccAddrFromMnemonic returns bech32 account address from mnemonic seeds
func GetBech32AccAddrFromMnemonic(mnemonic string, password string) (string, error) {
	seed := bip39.NewSeed(mnemonic, password)

	masterKey, ch := hd.ComputeMastersFromSeed(seed)
	privateKey, err := hd.DerivePrivateKeyForPath(masterKey, ch, hd.FullFundraiserPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to derive private key from hd path")
		return "", err
	}

	address, err := bech32.ConvertAndEncode(sdk.Bech32PrefixAccAddr, secp256k1.PrivKeySecp256k1(privateKey).PubKey().Address())
	if err != nil {
		log.Error().Err(err).Msg("failed to convert and encode to bech32 account address")
		return "", err
	}

	return address, nil
}

// GetPrivateKeyFromMnemonic returns a private key from mnemonic seeds
func GetPrivateKeyFromMnemonic(mnemonic string, password string) ([32]byte, error) {
	seed := bip39.NewSeed(mnemonic, password)
	masterKey, ch := hd.ComputeMastersFromSeed(seed)

	privateKey, err := hd.DerivePrivateKeyForPath(masterKey, ch, hd.FullFundraiserPath) // "44'/118'/0'/0/0"
	if err != nil {
		log.Error().Err(err).Msg("failed to derive private key from hd path")
		return privateKey, err
	}
	return privateKey, nil
}
