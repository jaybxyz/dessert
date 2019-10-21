package cmd

import (
	"github.com/spf13/cobra"
)

func getTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:     "tx",
		Aliases: []string{"t", "txs", "transaction", "transactions"},
		Short:   "Transaction subcommands",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return txCmd
}
