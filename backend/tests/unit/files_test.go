package app

import (
	"milonga/milonga/dbman"
	"os"
	"path/filepath"
	"testing"
)

func TestExitsFile(t *testing.T) {
	// Prueba con un archivo que existe
	tmpfile, err := os.CreateTemp("", "test-exists-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	if !dbman.ExitsFile(tmpfile.Name()) {
		t.Errorf("ExitsFile() returned false for existing file")
	}
	
	// Prueba con un archivo que no existe
	nonExistentFile := filepath.Join(os.TempDir(), "non-existent-file-test")
	_ = os.Remove(nonExistentFile) // Asegurarse de que no existe
	
	if dbman.ExitsFile(nonExistentFile) {
		t.Errorf("ExitsFile() returned true for non-existent file")
	}
}

func TestOpenFile(t *testing.T) {
	// Crear un archivo temporal con contenido conocido
	content := []byte("test content")
	tmpfile, err := os.CreateTemp("", "test-open-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	
	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	
	// Probar que podemos leer el archivo
	got := dbman.OpenFile(tmpfile.Name())
	if string(got) != string(content) {
		t.Errorf("OpenFile() = %v, want %v", string(got), string(content))
	}
	
	// Probar con un archivo que no existe
	nonExistentFile := filepath.Join(os.TempDir(), "non-existent-file-test")
	_ = os.Remove(nonExistentFile) // Asegurarse de que no existe
	
	got = dbman.OpenFile(nonExistentFile)
	if got != nil {
		t.Errorf("OpenFile() returned non-nil for non-existent file")
	}
}

func TestRemoveFile(t *testing.T) {
	// Crear un archivo temporal para eliminar
	tmpfile, err := os.CreateTemp("", "test-remove-*")
	if err != nil {
		t.Fatal(err)
	}
	tmpPath := tmpfile.Name()
	tmpfile.Close()
	
	// Verificar que existe antes de eliminar
	if !dbman.ExitsFile(tmpPath) {
		t.Fatalf("Test file does not exist before removal: %s", tmpPath)
	}
	
	// Eliminar el archivo
	if err := dbman.RemoveFile(tmpPath); err != nil {
		t.Errorf("RemoveFile() error = %v", err)
	}
	
	// Verificar que ya no existe
	if dbman.ExitsFile(tmpPath) {
		t.Errorf("File still exists after RemoveFile(): %s", tmpPath)
		// Limpiar para que no quede basura
		_ = os.Remove(tmpPath)
	}
	
	// Probar con un archivo que no existe
	nonExistentFile := filepath.Join(os.TempDir(), "non-existent-file-test")
	_ = os.Remove(nonExistentFile) // Asegurarse de que no existe
	
	if err := dbman.RemoveFile(nonExistentFile); err != nil {
		t.Errorf("RemoveFile() returned error for non-existent file: %v", err)
	}
}