package config

import (
	"fmt"
	"log"
	"os"
)

type ServerConfig struct {
	InterfaceName   string
	ServerPublicKey string
	ServerEndpoint  string
	ServerPort      int
}

func Load() ServerConfig {
	cfg := ServerConfig{
		InterfaceName:   os.Getenv("WG_INTERFACE"),
		ServerPublicKey: os.Getenv("WG_SERVER_PUBLIC_KEY"),
		ServerEndpoint:  os.Getenv("WG_SERVER_ENDPOINT"),
	}

	if cfg.InterfaceName == "" ||
		cfg.ServerPublicKey == "" ||
		cfg.ServerEndpoint == "" {
		log.Fatal("missing WireGuard server config")
	}

	// port default
	cfg.ServerPort = 51820
	if p := os.Getenv("WG_SERVER_PORT"); p != "" {
		fmt.Sscanf(p, "%d", &cfg.ServerPort)
	}

	return cfg
}
