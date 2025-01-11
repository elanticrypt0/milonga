package vigilante

import (
	"fmt"
	"milonga/internal/utils"
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

	// Validar la longitud de la clave de encriptación
	if len(config.EncryptionKey) != 32 {
		return nil, fmt.Errorf("encryption key must be exactly 32 bytes long")
	}

	return &config, nil
}

func GetPasswordTokenEncryptionKey() (string, error) {
	configFile := utils.GetAppRootPath() + "/config/app_config.toml"

	config, err := LoadPassTokenConfig(configFile)
	if err != nil {
		return "", fmt.Errorf("failed to load password token config: %w", err)
	}

	encryptionKey := config.EncryptionKey
	return encryptionKey, nil
}
