package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/paniccaaa/adamenblog/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: init config: cleanenv
	cfg, cfgDB := config.MustLoad()

	fmt.Println(cfg, "\n", cfgDB)

	//TODO: init logger: slog
	log := setupLogger(cfg.Env)
	
	log.Info(
		"starting adamen-blog", 
		slog.String("env", cfg.Env),
	)

	log.Debug("debug messages are enabled")
	
	//TODO: init storage postgresql

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
