package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	configPath := fmt.Sprint(dir, "/configs/", "config.ini")

	configLocationPtr := flag.String("path", configPath, "Path to the config file")

	flag.Parse()

	authConfig := loadConfig(configLocationPtr)
	tokenService := NewApiTokenService(authConfig)
	token := tokenService.GetToken()
	fmt.Println(token)
}
