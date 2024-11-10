package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/usysrc/obs/cmd"
)

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configDir := filepath.Join(home, ".config", "obsidian-cli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	cobra.OnInitialize(initConfig)

	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
