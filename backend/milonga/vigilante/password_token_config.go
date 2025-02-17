package vigilante

import (
	"fmt"
	"milonga/milonga/utils"
	"os"

	"github.com/BurntSushi/toml"
)

type PassTokenConfig struct {
	EncryptionKey string `toml:"PASSTOKEN_ENCRYPTIONKEY"`
}

// LoadPassTokenConfig lee la configuración del archivo TOML especificado
func LoadPassTokenConfig(filename string) (*PassTokenConfig, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file does not exist: %s", filename)
	}

	var config PassTokenConfig
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	// Decodificar y validar la clave
	key, err := DecodeEncryptionKey(config.EncryptionKey)
	if err != nil {
		return nil, fmt.Errorf("invalid encryption key: %w", err)
	}

	// Guardar la clave decodificada
	config.EncryptionKey = string(key)

	return &config, nil
}

func GetPasswordTokenEncryptionKey() (string, error) {
	configFile := utils.GetAppRootPath() + "/config/app_config.toml"

	config, err := LoadPassTokenConfig(configFile)
	if err != nil {
		return "", fmt.Errorf("failed to load password token config: %w", err)
	}

	return config.EncryptionKey, nil
}
