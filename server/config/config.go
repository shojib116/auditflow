package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort       int
	FrontendDomain string
}

var (
	config     *Config
	configOnce sync.Once
)

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	httpPortStr := envOrExit("HTTP_PORT")
	port, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("HTTP_PORT must be a valid integer:", err)
		os.Exit(1)
	}

	if port < 1 || port > 65535 {
		fmt.Println("HTTP_PORT must be between 1 and 65535")
		os.Exit(1)
	}

	config = &Config{
		HttpPort:       port,
		FrontendDomain: envOrExit("FRONTEND_DOMAIN"),
	}
}

func envOrExit(envKey string) string {
	envVar := os.Getenv(envKey)
	if envVar == "" {
		fmt.Println(envKey, "environment varaible is required")
		os.Exit(1)
	}

	return envVar
}

func GetConfig() *Config {
	configOnce.Do(loadConfig)
	return config
}
