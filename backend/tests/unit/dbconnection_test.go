package app

import (
	"fmt"
	"milonga/milonga/dbman"
	"testing"

	"gorm.io/gorm"
)

func TestDBConnection_IsOk(t *testing.T) {
	tests := []struct {
		name   string
		conn   dbman.DBConnection
		want   bool
	}{
		{
			name: "connection without errors",
			conn: dbman.DBConnection{
				DBConfig: dbman.NewDBConfig("test", "sqlite", "", "", "", "", ":memory:"),
				Instance: nil,
				ErrConn:  nil,
			},
			want: true,
		},
		{
			name: "connection with errors",
			conn: dbman.DBConnection{
				DBConfig: dbman.NewDBConfig("test", "sqlite", "", "", "", "", ":memory:"),
				Instance: nil,
				ErrConn:  fmt.Errorf("error de conexion"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.conn.IsOk(); got != tt.want {
				t.Errorf("DBConnection.IsOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_IsConnected(t *testing.T) {
	tests := []struct {
		name   string
		conn   dbman.DBConnection
		want   bool
	}{
		{
			name: "connection established",
			conn: dbman.DBConnection{
				DBConfig: dbman.NewDBConfig("test", "sqlite", "", "", "", "", ":memory:"),
				Instance: &gorm.DB{},
				ErrConn:  nil,
			},
			want: true,
		},
		{
			name: "connection not established",
			conn: dbman.DBConnection{
				DBConfig: dbman.NewDBConfig("test", "sqlite", "", "", "", "", ":memory:"),
				Instance: nil,
				ErrConn:  nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.conn.IsConnected(); got != tt.want {
				t.Errorf("DBConnection.IsConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test de integración para la conexión SQLite usando archivo
func TestDBConnection_Connect(t *testing.T) {
	// Usar archivo temporal en lugar de :memory: para evitar problemas de paths
	dbConfig := dbman.NewDBConfig("test", "sqlite", "", "", "", "", "test_db.db")
	
	conn := dbman.NewDBConn(dbConfig)
	err := conn.Connect(".")
	
	if err != nil {
		t.Errorf("DBConnection.Connect() error = %v", err)
	}
	
	if !conn.IsConnected() {
		t.Errorf("DBConnection.Connect() did not establish connection")
	}
	
	// Limpiar el archivo de base de datos temporal
	conn.Instance = nil
	// Intentamos eliminar el archivo, pero no es crítico si falla
	_ = dbman.RemoveFile("./test_db.db")
	_ = dbman.RemoveFile("test_db.db")
}