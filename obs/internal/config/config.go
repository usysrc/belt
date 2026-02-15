package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetVault(v *viper.Viper) (string, error) {
	vault := v.GetString("vault")
	if vault == "" {
		return "", fmt.Errorf("vault not configured. Run 'obs config --vault \"Your Vault\"' first")
	}
	return vault, nil
}

func GetTargetFolder(v *viper.Viper) (string, error) {
	targetFolder := v.GetString("targetFolder")
	return targetFolder, nil
}
