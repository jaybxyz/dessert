package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "initialize config.toml with required configuration params",
	Long:    `initialize config.toml with required configuration params`,
	RunE:    initCmdHandler,
}

func initCmdHandler(cmd *cobra.Command, args []string) error {
	filePath, _ := filepath.Abs("./" + cfgFile)

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create configuration file: %s", err)
	}

	f.WriteString(`[client]
rpc_endpoint = "<ip>:<port>"
lcd_endpoint = "<ip>:<port>"

[wallet]
network = "<network>"
mnemonic = "<mnemonic phrases>"
`)
	return nil
}
