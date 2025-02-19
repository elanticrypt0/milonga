package app

import (
	"milonga/milonga/dbman"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewDBConn(t *testing.T) {
	type args struct {
		connData dbman.DBConfig
	}
	tests := []struct {
		name string
		args args
		want dbman.DBConnection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dbman.NewDBConn(tt.args.connData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDBConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_IsOk(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			if got := me.IsOk(); got != tt.want {
				t.Errorf("DBConnection.IsOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_IsConnected(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			if got := me.IsConnected(); got != tt.want {
				t.Errorf("DBConnection.IsConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_Connect(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	type args struct {
		rootPath string
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
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			if err := me.Connect(tt.args.rootPath); (err != nil) != tt.wantErr {
				t.Errorf("DBConnection.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBConnection_connect2Mysql(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	tests := []struct {
		name    string
		fields  fields
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			got, err := me.connect2Mysql()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBConnection.connect2Mysql() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnection.connect2Mysql() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_connect2Postgres(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	tests := []struct {
		name    string
		fields  fields
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			got, err := me.connect2Postgres()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBConnection.connect2Postgres() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnection.connect2Postgres() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_connect2SQLite(t *testing.T) {
	type fields struct {
		DBConfig DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	type args struct {
		rootPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			got, err := me.connect2SQLite(tt.args.rootPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBConnection.connect2SQLite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnection.connect2SQLite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnection_returnErrConn(t *testing.T) {
	type fields struct {
		DBConfig dbman.DBConfig
		Instance *gorm.DB
		ErrConn  error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBConnection{
				DBConfig: tt.fields.DBConfig,
				Instance: tt.fields.Instance,
				ErrConn:  tt.fields.ErrConn,
			}
			if err := me.returnErrConn(); (err != nil) != tt.wantErr {
				t.Errorf("DBConnection.returnErrConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
