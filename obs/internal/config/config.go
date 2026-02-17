package config

import (
	"fmt"
	"os"

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

func GetVaultPath(v *viper.Viper) (string, error) {
	if envVaultPath := os.Getenv("OBS_VAULT_PATH"); envVaultPath != "" {
		return envVaultPath, nil
	}

	vaultPath := v.GetString("vaultPath")
	if vaultPath == "" {
		return "", fmt.Errorf(
			"vault path not configured. Run 'obs config --vaultPath \"/path/to/vault\"' or set OBS_VAULT_PATH",
		)
	}

	return vaultPath, nil
}
