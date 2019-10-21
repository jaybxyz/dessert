package cmd

import (
	"fmt"
	"os"

	"github.com/kogisin/dessert/client"
	"github.com/kogisin/dessert/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

const (
	cfgFile      = "./config.toml"
	logLevelJSON = "json"
	logLevelText = "text"
)

var (
	cfgFilePath string
	logLevel    string
	logFormat   string
)

var rootCmd = &cobra.Command{
	Use:     "dessert [command]",
	Short:   "dessert is a sample project that provides various examples of how to work with Cosmos SDK based projects.",
	Long:    `dessert is a sample project that provides various examples of how to work with Cosmos SDK based projects.`,
	PreRunE: cmdHandler,
}

func init() {
	// Available Commands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(walletCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(txCmd)

	// Flags
	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config-path", cfgFile, "config.toml file path")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", zerolog.InfoLevel.String(), "logging level")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", logLevelJSON, "logging format; must be either json or text")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdHandler(cmd *cobra.Command, args []string) error {
	logLvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(logLvl)

	switch logFormat {
	case logLevelJSON:
		// JSON is the default logging format

	case logLevelText:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	default:
		return fmt.Errorf("invalid logging format: %s", logFormat)
	}

	cfg := config.ParseConfig(cfgFile)

	cw, err := client.New(cfg.Client.RPCEndpoint, cfg.Client.LCDEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to start RPC client")
	}

	defer cw.Stop() // nolint: errcheck

	return nil
}
