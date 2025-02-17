package tests

import (
	"dbman"
	"os"
	"testing"
)

func TestSetRootPath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Valid absolute path",
			path:    "/var/lib/data",
			wantErr: false,
		},
		{
			name:    "Empty path",
			path:    "",
			wantErr: true,
		},
		{
			name:    "Relative path",
			path:    "data/db",
			wantErr: true,
		},
		{
			name:    "Path with parent directory reference",
			path:    "/var/lib/../data",
			wantErr: true,
		},
		{
			name:    "Path with special characters",
			path:    "/var/lib/data?test",
			wantErr: true,
		},
		{
			name:    "Path with home directory reference",
			path:    "~/data",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := dbman.New()
			err := db.SetRootPath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRootPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadConfigEnv(t *testing.T) {
	// Setup test environment
	tempEnvFile := ".env.test"
	defer os.Remove(tempEnvFile)

	tests := []struct {
		name     string
		envVars  map[string]string
		wantErr  bool
		setupEnv func()
	}{
		{
			name: "Valid configuration",
			envVars: map[string]string{
				"DB_CONN_NAME": "testdb",
				"DB_ENGINE":    "postgres",
				"DB_HOST":      "localhost",
				"DB_PORT":      "5432",
				"DB_USER":      "test",
				"DB_PASSWORD":  "test",
				"DB_NAME":      "testdb",
			},
			wantErr: false,
		},
		{
			name: "Missing required variable",
			envVars: map[string]string{
				"DB_CONN_NAME": "testdb",
				// Missing DB_ENGINE
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "test",
				"DB_PASSWORD": "test",
				"DB_NAME":     "testdb",
			},
			wantErr: true,
		},
		{
			name: "Invalid port number",
			envVars: map[string]string{
				"DB_CONN_NAME": "testdb",
				"DB_ENGINE":    "postgres",
				"DB_HOST":      "localhost",
				"DB_PORT":      "invalid",
				"DB_USER":      "test",
				"DB_PASSWORD":  "test",
				"DB_NAME":      "testdb",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary .env file
			f, err := os.Create(tempEnvFile)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			// Write environment variables
			for k, v := range tt.envVars {
				if _, err := f.WriteString(k + "=" + v + "\n"); err != nil {
					t.Fatal(err)
				}
			}

			db := dbman.New()
			err = db.LoadConfigEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
