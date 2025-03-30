package app

import (
	"milonga/milonga/dbman"
	"os"
	"path/filepath"
	"testing"
)

// TestDBConfig devuelve una configuración de base de datos SQLite para pruebas
func TestDBConfig(t *testing.T) dbman.DBConfig {
	return dbman.NewDBConfig("test", "sqlite", "", "", "", "", ":memory:")
}

// TestDBConfigFile devuelve una configuración de base de datos SQLite basada en archivos para pruebas
func TestDBConfigFile(t *testing.T) dbman.DBConfig {
	// Crear un directorio temporal para las pruebas
	tempDir := os.TempDir()
	dbPath := filepath.Join(tempDir, "test_milonga.db")
	
	// Eliminar el archivo si ya existe
	_ = os.Remove(dbPath)
	
	return dbman.NewDBConfig("test_file", "sqlite", "", "", "", "", dbPath)
}

// SetupTestDB configura una base de datos SQLite en memoria para pruebas
func SetupTestDB(t *testing.T) *dbman.DBConnection {
	config := TestDBConfig(t)
	dbConn := dbman.NewDBConn(config)
	err := dbConn.Connect("")
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	
	if !dbConn.IsConnected() {
		t.Fatal("Database connection failed")
	}
	
	return &dbConn
}

// SetupTestDBFile configura una base de datos SQLite basada en archivos para pruebas
func SetupTestDBFile(t *testing.T) *dbman.DBConnection {
	config := TestDBConfigFile(t)
	dbConn := dbman.NewDBConn(config)
	err := dbConn.Connect("")
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	
	if !dbConn.IsConnected() {
		t.Fatal("Database connection failed")
	}
	
	return &dbConn
}

// CleanupTestDB limpia los recursos de la base de datos de prueba
func CleanupTestDB(t *testing.T, conn *dbman.DBConnection) {
	// Para SQLite en memoria, no necesitamos hacer nada más
	// La base de datos se elimina cuando se cierra la conexión
	
	// Si estamos usando un archivo SQLite, podríamos eliminarlo aquí
	if conn != nil && conn.DBConfig.DBName != ":memory:" {
		// Cerrar la conexión y eliminar el archivo
		if conn.IsConnected() {
			_ = os.Remove(conn.DBConfig.DBName)
		}
	}
}