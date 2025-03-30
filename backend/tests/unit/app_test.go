package app

import (
	"context"
	"milonga/milonga/app"
	"milonga/milonga/dbman"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestApp_GetCurrentEnviroment(t *testing.T) {
	type fields struct {
		Server *fiber.App
		Config *app.Config
		DB     *dbman.DBMan
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "deve environment",
			fields: fields{
				Server: nil,
				Config: &app.Config{
					Name:       "Test App",
					Version:    "1.0.0",
					Enviroment: "dev",
				},
				DB: nil,
			},
			want: "dev",
		},
		{
			name: "production environment",
			fields: fields{
				Server: nil,
				Config: &app.Config{
					Name:       "Test App",
					Version:    "1.0.0",
					Enviroment: "production",
				},
				DB: nil,
			},
			want: "production",
		},
		{
			name: "testing environment",
			fields: fields{
				Server: nil,
				Config: &app.Config{
					Name:       "Test App",
					Version:    "1.0.0",
					Enviroment: "testing",
				},
				DB: nil,
			},
			want: "testing",
		},
		{
			name: "empty environment",
			fields: fields{
				Server: nil,
				Config: &app.Config{
					Name:       "Test App",
					Version:    "1.0.0",
					Enviroment: "",
				},
				DB: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &app.App{
				Server: tt.fields.Server,
				Config: tt.fields.Config,
				DB:     tt.fields.DB,
			}
			if got := me.GetCurrentEnviroment(); got != tt.want {
				t.Errorf("App.GetCurrentEnviroment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		wantApp    bool
	}{
		{
			name:       "valid config path",
			configPath: "../../config/app_config.toml",
			wantApp:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := app.New(tt.configPath)
			if (got != nil) != tt.wantApp {
				t.Errorf("New() = %v, want %v", got, tt.wantApp)
			}
			if got != nil && got.Config == nil {
				t.Errorf("New() Config is nil, expected non-nil Config")
			}
		})
	}
}

func TestApp_SetCtx(t *testing.T) {
	testApp := &app.App{
		Config: &app.Config{
			Name:       "Test App",
			Version:    "1.0.0",
			Enviroment: "testing",
		},
	}
	
	ctx := context.Background()
	testApp.SetCtx(ctx)
	
	if testApp.Ctx == nil {
		t.Errorf("SetCtx() failed to set context, Ctx is nil")
	}
}

func TestApp_ConsoleMessage(t *testing.T) {
	testApp := &app.App{
		Config: &app.Config{
			Name:       "Test App",
			Version:    "1.0.0",
			Enviroment: "testing",
		},
	}
	
	// This just tests that the method doesn't panic
	testApp.ConsoleMessage("Test message")
}
