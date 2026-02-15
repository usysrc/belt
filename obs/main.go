package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"codeberg.org/usysrc/belt/obs/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	configDir := filepath.Join(home, ".config", "obsidian-cli")
	if err := os.MkdirAll(configDir, 0750); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			// Config file not found; ignore error if desired
			log.Printf("Config file not found in %s, using defaults", configDir)
		} else {
			// Config file was found but another error was produced
			log.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	cobra.OnInitialize(initConfig)

	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
