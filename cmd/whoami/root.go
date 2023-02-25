package cmd

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/docker/go-units"
	"github.com/pprishchepa/whoami/internal/random"
	"github.com/pprishchepa/whoami/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

var debugMode bool

var rootCmd = &cobra.Command{
	Short: "Whoami is blazing-fast testing upstream.",
	Use:   "whoami",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
			Level(zerolog.TraceLevel).With().Timestamp().Logger()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		rand.Seed(time.Now().UnixNano())
		random.Randomize(5 * units.MiB)
		opts := server.DefaultOptions
		opts.Debug = debugMode
		return server.Serve(cmd.Context(), opts)
	},
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("WHOAMI")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "D", viper.GetBool("DEBUG"), "Enable debug mode")
}

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		select {
		case <-ch:
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
