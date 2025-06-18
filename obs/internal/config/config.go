package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetVault() (string, error) {
	vault := viper.GetString("vault")
	if vault == "" {
		return "", fmt.Errorf("vault not configured. Run 'obs config --vault \"Your Vault\"' first")
	}
	return vault, nil
}

func GetTargetFolder() (string, error) {
	targetFolder := viper.GetString("targetFolder")
	return targetFolder, nil
}
