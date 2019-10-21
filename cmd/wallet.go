package cmd

import (
	"github.com/spf13/cobra"
)

func getWalletCmd() *cobra.Command {
	walletCmd := &cobra.Command{
		Use:     "wallet",
		Aliases: []string{"w"},
		Short:   "Wallet subcommands",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return walletCmd
}
