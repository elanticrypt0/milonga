package app

import (
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
		// TODO: Add test cases.
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
