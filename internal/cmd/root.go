package cmd

import (
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logger *slog.Logger

var rootCmd = &cobra.Command{
	Short: "Micro finance service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Panic("Failed to start application", err)
	}
}

func initApp() {
	logger = slog.NewStdLogger().Logger
	logger.Info("Application starting up")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Errorf("Fatal error config file: %w \n", err)
	}
}
