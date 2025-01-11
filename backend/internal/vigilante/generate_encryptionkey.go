package vigilante

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// GenerateEncryptionKey genera una clave segura de 32 bytes
func GenerateEncryptionKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("error generating key: %w", err)
	}

	// Puedes elegir entre estas dos formas de codificar la clave:

	// Opción 1: Hex encoding (resultará en 64 caracteres)
	// return hex.EncodeToString(key), nil

	// Opción 2: Base64 encoding (resultará en aproximadamente 44 caracteres)
	return base64.StdEncoding.EncodeToString(key), nil
}

// DecodeEncryptionKey decodifica una clave en formato hex o base64 a bytes
func DecodeEncryptionKey(encodedKey string) ([]byte, error) {
	// Intentar primero como hex
	if key, err := hex.DecodeString(encodedKey); err == nil && len(key) == 32 {
		return key, nil
	}

	// Intentar como base64
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil || len(key) != 32 {
		return nil, fmt.Errorf("invalid key format or length")
	}

	return key, nil
}
