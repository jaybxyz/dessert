package cmd

import (
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:     "wallet",
	Aliases: []string{"w"},
	Short:   "wallet subcommands",
	Long:    `wallet subcommands`,
	RunE:    walletCmdHandler,
}

func walletCmdHandler(cmd *cobra.Command, args []string) error {

	return nil
}
