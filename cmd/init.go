package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func getInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "initialize config.toml with required configuration params",
		RunE: func(cmd *cobra.Command, args []string) error {

			filePath, _ := filepath.Abs("./" + cfgFile)

			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("unable to create configuration file: %s", err)
			}

			f.WriteString(`[client]
rpc_endpoint = "<ip>:<port>"
lcd_endpoint = "<ip>:<port>"

[wallet]
mnemonic = "<mnemonic phrases>"
`)
			f.Sync()

			return nil
		},
	}
	return initCmd
}
