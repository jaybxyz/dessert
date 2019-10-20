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
	logLevelJSON = "json"
	logLevelText = "text"
)

var (
	logLevel  string
	logFormat string
)

var rootCmd = &cobra.Command{
	Use:   "dessert [config.toml]",
	Args:  cobra.ExactArgs(1),
	Short: "dessert is a sample project that provides delicious desserts for Cosmos SDK based projects",
	Long:  `A Sample Project that provides various examples of how to work with Cosmos SDK Based Project.`,
	RunE:  cmdHandler,
}

func init() {
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

	cfgFile := args[0]
	cfg := config.ParseConfig(cfgFile)

	cw, err := client.New(cfg.RPCEndpoint, cfg.LCDEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to start RPC client")
	}

	defer cw.Stop() // nolint: errcheck

	return nil
}
