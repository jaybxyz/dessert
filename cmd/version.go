package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the app version",
	Long:  `Print the app version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
