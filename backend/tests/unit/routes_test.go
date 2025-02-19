package app

import (
	"milonga/api/handlers"
	"milonga/milonga/app"
	"testing"

	"gorm.io/gorm"
)

func TestActivateRoutes(t *testing.T) {
	type args struct {
		app       *app.App
		services  []handlers.CriticalService
		dbs2check []*gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers.ActivateRoutes(tt.args.app, tt.args.services, tt.args.dbs2check)
		})
	}
}
