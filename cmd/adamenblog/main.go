package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/paniccaaa/adamenblog/internal/config"
	"github.com/paniccaaa/adamenblog/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: init config: cleanenv
	cfg, cfgDB := config.MustLoad()

	//TODO: init logger: slog
	log := setupLogger(cfg.Env)

	log.Info(
		"starting adamenblog",
		slog.String("env", cfg.Env),
	)

	log.Debug("debug messages are enabled")

	//TODO: init storage postgresql
	storage, err := postgres.NewPostgres(cfgDB)
	if err != nil {
		log.Error("failed to init storage: %s", err)
		os.Exit(1)
	}

	fmt.Println(storage)

	//TODO: init router chi, chi-render

	//TODO: run server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default: //if env config is invalid, set prod settings by default due to security
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
