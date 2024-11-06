package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	Version     string `yaml:"version" env-default:"0.0.0"`
	Server      `yaml:"server"`
	Storage     `yaml:"storage"`
	CacheServer `yaml:"cache_server"`
}

type Server struct {
	Address         string        `yaml:"address" env-default:"localhost:8080"`
	Timeout         time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout     time.Duration `yaml:"idle_timeout" env-default:"30s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-default:"10s"`
}

type Storage struct {
	Server   string `yaml:"server" env-default:"localhost"`
	Database string `yaml:"database" env-default:"postgres"`
	Port     int    `yaml:"port" env-default:"5432"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env:"storage_password" env-required:"true"`
}

type CacheServer struct {
	Address  string `yaml:"address" env-default:"localhost:6379"`
	Password string `yaml:"password" env:"cache_server_password" env-required:"true"`
	Database int    `yaml:"database" env-default:"0"`
}

func New() *Config {
	// Get the path to the configuration file from the environment
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Fprintf(os.Stderr, "Failed to find CONFIG_PATH environment variable")
		os.Exit(1)
	}

	// Check if the configuration file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Failed to find configuration file in %s", configPath)
		os.Exit(1)
	}

	// Read the configuration file
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read configuration: %s", err.Error())
		os.Exit(1)
	}

	return &cfg
}
