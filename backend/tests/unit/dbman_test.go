package app

import (
	"milonga/milonga/dbman"
	"testing"
)

func TestDBManNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create new DBMan instance",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dbman.New()
			if got == nil {
				t.Errorf("New() returned nil")
			}
		})
	}
}

func TestDBMan_SetRootPath(t *testing.T) {
	type args struct {
		rootpath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set valid root path",
			args: args{
				rootpath: "/tmp",
			},
			wantErr: false,
		},
		{
			name: "Set empty root path",
			args: args{
				rootpath: "",
			},
			wantErr: true,
		},
		{
			name: "Set relative path",
			args: args{
				rootpath: "relative/path",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := dbman.New()
			if err := me.SetRootPath(tt.args.rootpath); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.SetRootPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_LoadConfigToml(t *testing.T) {
	tests := []struct {
		name    string
		path    string
	}{
		{
			name:    "Load test config",
			path:    "../../config/test_db_config.toml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := dbman.New()
			// La función no retorna valor, solo comprobamos que no cause pánico
			me.LoadConfigToml(tt.path)
		})
	}
}