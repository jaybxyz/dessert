package cmd

import (
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"q"},
	Short:   "querying subcommands",
	Long:    `querying subcommands`,
	RunE:    queryCmdHandler,
}

func queryCmdHandler(cmd *cobra.Command, args []string) error {

	return nil
}
