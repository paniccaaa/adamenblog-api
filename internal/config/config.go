package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type ConfigDatabase struct {
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"USER_NAME"`
	Password string `env:"PASSWORD"`
}

func MustLoad() (*Config, *ConfigDatabase) {
	configPath := "./config/local.yaml"

	//check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	if err := LoadEnv(); err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	var cfgDB ConfigDatabase
	err := cleanenv.ReadEnv(&cfgDB)
	if err != nil {
		log.Fatalf("cannot read .env db: %s", err)
	}
	
	return &cfg, &cfgDB
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
