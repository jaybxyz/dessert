package cmd

import (
	"fmt"
	"os"

	"github.com/kogisin/dessert/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
	Use:   "seasoning [config-file]",
	Args:  cobra.ExactArgs(1),
	Short: "Seasoning is like salt & pepper for Cosmos SDK based projects",
	Long: `A cosmos Hub data aggregator. It improves the Hub's data accessibility
by providing an indexed PostgreSQL database exposing aggregated resources and
models such as blocks, validators, pre-commits, transactions, and various aspects
of the governance module. Juno is meant to run with a GraphQL layer on top so that
it even further eases the ability for developers and downstream clients to answer
queries such as "what is the average gas cost of a block?" while also allowing
them to compose more aggregate and complex queries.`,
	RunE: cmdHandler,
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

	cp, err := client.New(cfg.RPCEndpoint, cfg.LCDEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to start RPC client")
	}

	defer cp.Stop() // nolint: errcheck

	return nil
}