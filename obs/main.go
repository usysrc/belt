package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"codeberg.org/usysrc/belt/obs/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig(cfg *viper.Viper) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configDir := filepath.Join(home, ".config", "obsidian-cli")
	if err := os.MkdirAll(configDir, 0750); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg.AddConfigPath(configDir)
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AutomaticEnv()

	if err := cfg.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			// Config file not found; ignore error if desired
			fmt.Printf("Config file not found in %s\n", configDir)
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	cfg := viper.New()

	cobra.OnInitialize(func() { initConfig(cfg) })

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.NewRootCmd(home, cfg).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
