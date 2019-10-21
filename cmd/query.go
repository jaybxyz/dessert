package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/rs/zerolog/log"

	dessertclient "github.com/kogisin/dessert/client"
	"github.com/kogisin/dessert/config"
	"github.com/kogisin/dessert/models"
	"github.com/kogisin/dessert/wallet"

	"github.com/spf13/cobra"

	resty "gopkg.in/resty.v1"
)

/*
	Zerolog sample
		log.Error().Err(err).Str("hash", tx.TxHash).Msg("failed to persist transaction")
*/

func getQueryCmd(cfg config.Config, cw dessertclient.ConfigWrapper) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query [command]",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		balanceCmd(cfg, cw),
		client.LineBreak,
	)

	return queryCmd
}

func balanceCmd(cfg config.Config, cw dessertclient.ConfigWrapper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Get your account balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get bech32 account address from mnemonic
			cosmosAddress, _ := wallet.GetBech32AccAddrFromMnemonic(cfg.Wallet.Mnemonic, "")

			// Query account information
			resp, err := resty.R().Get(cfg.Client.LCDEndpoint + "/auth/accounts/" + cosmosAddress)
			if err != nil {
				log.Error().Err(err).Str(cfg.Client.LCDEndpoint, "/auth/accounts/"+cosmosAddress).Msg("failed to fetch data")
				return err
			}

			var accountResp models.Account
			err = json.Unmarshal(resp.Body(), &accountResp)
			if err != nil {
				log.Error().Err(err).Msg("failed to unmarshal models.Account data")
				return err
			}

			fmt.Println("Address: ", accountResp.Value.Address)
			fmt.Println("Denom: ", accountResp.Value.Coins[0].Denom)
			fmt.Println("Amount: ", accountResp.Value.Coins[0].Amount)
			fmt.Println("Address Number: ", accountResp.Value.AccountNumber)
			fmt.Println("Sequence: ", accountResp.Value.Sequence)

			return nil
		},
	}

	return cmd
}
