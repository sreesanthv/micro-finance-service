package cmd

import (
	"context"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(tryCmd)
}

var tryCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API",
	Run: func(cmd *cobra.Command, args []string) {
		initApp()
		startApi()
	},
}

func startApi() {
	dbPool, err := pgxpool.Connect(context.Background(), viper.GetString("db-url"))
	if err != nil {
		logger.Panic("Error in connecting DB:", err)
	}
	defer dbPool.Close()
	logger.Info("Connected to DB. Max conn pool count:", dbPool.Config().MaxConns)

	redis := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis-url"),
		Password: "",
		DB:       0,
	})

	http.ListenAndServe(":3000", Routes(dbPool, redis))
}
