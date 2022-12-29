package main

import (
	"fmt"
	"project/config"
	"project/config/keys"
	"sync"

	"github.com/spf13/viper"
)

func main() {
	// Set configuration file type and name.
	viper.SetDefault(keys.AppName, "BTPN Syariah Fullstack Developer")
	viper.SetDefault(keys.BaseURL, "localhost:8008")
	viper.SetDefault(keys.DatabaseURL, "postgres://postgres:postgres@localhost:5432/vip_btpns")
	viper.SetDefault(keys.DatabaseDriver, "postgres")
	viper.SetDefault(keys.Environment, "development")
	viper.SetDefault(keys.HostAddress, "localhost")
	viper.SetDefault(keys.HostPort, "8008")

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")
	if readErr := viper.ReadInConfig(); readErr != nil {
		// Write config file if it doesn't exist
		if _, ok := readErr.(viper.ConfigFileNotFoundError); ok {
			if writeErr := viper.SafeWriteConfig(); writeErr != nil {
				fmt.Printf("Failed to create config file(s): %s\n", writeErr)
			}
		}
		fmt.Printf("Failed to load config file: %s . Using defaults\n", readErr)
	}

	config := config.NewConfig()
	server := InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
