package app

import (
	"milonga/milonga/app"
	"milonga/milonga/healthcheck"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TestNewHealthCheck(t *testing.T) {
	type args struct {
		app      *app.App
		services []healthcheck.CriticalService
		dbs      []*gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *HealthCheck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := healthcheck.NewHealthCheck(tt.args.app, tt.args.services, tt.args.dbs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthCheck_CheckDBConnection(t *testing.T) {
	type fields struct {
		app            *app.App
		Services2Check []healthcheck.CriticalService
		DB2Check       []*gorm.DB
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &healthcheck.HealthCheck{
				app:            tt.fields.app,
				Services2Check: tt.fields.Services2Check,
				DB2Check:       tt.fields.DB2Check,
			}
			if err := me.CheckDBConnection(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck.CheckDBConnection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHealthCheck_checkCriticalServices(t *testing.T) {
	type fields struct {
		app            *app.App
		Services2Check []healthcheck.CriticalService
		DB2Check       []*gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   healthcheck.ServiceStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &healthcheck.HealthCheck{
				app:            tt.fields.app,
				Services2Check: tt.fields.Services2Check,
				DB2Check:       tt.fields.DB2Check,
			}
			if got := me.checkCriticalServices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HealthCheck.checkCriticalServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthCheck_Check(t *testing.T) {
	type fields struct {
		app            *app.App
		Services2Check []healthcheck.CriticalService
		DB2Check       []*gorm.DB
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &healthcheck.HealthCheck{
				app:            tt.fields.app,
				Services2Check: tt.fields.Services2Check,
				DB2Check:       tt.fields.DB2Check,
			}
			if err := me.Check(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
