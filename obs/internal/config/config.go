package config

import (
	"fmt"
	"os"
)

type Getter interface {
	GetString(key string) string
}

func GetVault(g Getter) (string, error) {
	vault := g.GetString("vault")
	if vault == "" {
		return "", fmt.Errorf("vault not configured. Run 'obs config --vault \"Your Vault\"' first")
	}

	return vault, nil
}

func GetTargetFolder(g Getter) (string, error) {
	targetFolder := g.GetString("targetFolder")

	return targetFolder, nil
}

func GetVaultPath(g Getter) (string, error) {
	if envVaultPath := os.Getenv("OBS_VAULT_PATH"); envVaultPath != "" {
		return envVaultPath, nil
	}

	vaultPath := g.GetString("vaultPath")
	if vaultPath == "" {
		return "", fmt.Errorf(
			"vault path not configured. Run 'obs config --vaultPath \"/path/to/vault\"' or set OBS_VAULT_PATH",
		)
	}

	return vaultPath, nil
}
