package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)


type AuthConfig struct {
	clientSecret string
	clientId	 string
	tokenUrl	 string
}

func loadConfig(configLocation *string) AuthConfig {
	config, error := ini.Load(*configLocation)
	if error != nil {
		fmt.Println("Failed to read config file: %v", error)
		os.Exit(1)
	}

	authSection := config.Section("AUTH")

	authConfig := AuthConfig{
		authSection.Key("CLIENT_SECRET").String(),
		authSection.Key("CLIENT_ID").String(),
		authSection.Key("TOKEN_URL").String(),
	}

	return authConfig
}
