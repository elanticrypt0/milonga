package app

import (
	"milonga/milonga/dbman"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *dbman.DBMan
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dbman.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBMan_SetRootPath(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		rootpath string
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.SetRootPath(tt.args.rootpath); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.SetRootPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_LoadConfigToml(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.LoadConfigToml(tt.args.filepath)
		})
	}
}

func TestDBMan_LoadConfigEnv(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.LoadConfigEnv(); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.LoadConfigEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_addConn(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		connData DBConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.addConn(tt.args.connData)
		})
	}
}

func TestDBMan_addActiveConn(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		conn string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.addActiveConn(tt.args.conn)
		})
	}
}

func TestDBMan_GetInstance(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			got, err := me.GetInstance(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBMan.GetInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBMan.GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBMan_getInstanceIfExists(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dbman.DBConnection
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			got, err := me.getInstanceIfExists(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBMan.getInstanceIfExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBMan.getInstanceIfExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBMan_IsDBOk(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		connName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if got := me.IsDBOk(tt.args.connName); got != tt.want {
				t.Errorf("DBMan.IsDBOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBMan_Connect(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.Connect(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_SetPrimary(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.SetPrimary(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.SetPrimary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_SetSecondary(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
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
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.SetSecondary(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.SetSecondary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_SetSecurity(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	type args struct {
		name string
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
			me := &dbman.BMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			if err := me.SetSecurity(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DBMan.SetSecurity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBMan_PrintConnectionsList(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.PrintConnectionsList()
		})
	}
}

func TestDBMan_PrintActiveConnectionsList(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.PrintActiveConnectionsList()
		})
	}
}

func TestDBMan_CheckDefaultConnections(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			me.CheckDefaultConnections()
		})
	}
}

func TestDBMan_GetActiveConnectionsInstances(t *testing.T) {
	type fields struct {
		rootPath         string
		configPath       string
		connection       map[string]dbman.DBConnection
		activeConnection []string
		Primary          *gorm.DB
		Secondary        *gorm.DB
		Security         *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &dbman.DBMan{
				rootPath:         tt.fields.rootPath,
				configPath:       tt.fields.configPath,
				connection:       tt.fields.connection,
				activeConnection: tt.fields.activeConnection,
				Primary:          tt.fields.Primary,
				Secondary:        tt.fields.Secondary,
				Security:         tt.fields.Security,
			}
			got, err := me.GetActiveConnectionsInstances()
			if (err != nil) != tt.wantErr {
				t.Errorf("DBMan.GetActiveConnectionsInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBMan.GetActiveConnectionsInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}
