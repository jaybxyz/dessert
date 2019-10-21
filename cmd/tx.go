package cmd

import (
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"t", "txs", "transaction", "transactions"},
	Short:   "Transaction subcommands",
	Long:    `Transaction subcommands`,
	RunE:    txCmdHandler,
}

func txCmdHandler(cmd *cobra.Command, args []string) error {

	return nil
}
