package app

import (
	"milonga/milonga/app"
	"testing"
)

type TestConfig struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

func TestLoadConfig(t *testing.T) {
	// Esta prueba se limita a verificar que la función no causa pánico
	// y devuelve un objeto App no nulo
	got := app.LoadConfig("../../config/app_config.toml")
	if got == nil {
		t.Errorf("LoadConfig() returned nil")
	}
}

func TestLoadConfigFileNotExists(t *testing.T) {
	// Intentamos cargar un archivo que no existe
	// Lo esperado es que no cause pánico
	// La implementación actual puede devolver un objeto o nil
	// Lo importante es que no cause pánico
	app.LoadConfig("nonexistent_file.toml")
	// Test exitoso si llegamos aquí sin pánico
}