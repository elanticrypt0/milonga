package app

import (
	"context"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		scriptsPath string
		storagePath string
	}
	tests := []struct {
		name string
		args args
		want *RunPy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runpy.New(tt.args.scriptsPath, tt.args.storagePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunPy_runScript(t *testing.T) {
	type fields struct {
		runtime     string
		scriptsPath string
		storagePath string
	}
	type args struct {
		ctx            context.Context
		script2execute string
		filePathArg    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *runpy.ScriptResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &runpy.RunPy{
				runtime:     tt.fields.runtime,
				scriptsPath: tt.fields.scriptsPath,
				storagePath: tt.fields.storagePath,
			}
			got, err := me.runScript(tt.args.ctx, tt.args.script2execute, tt.args.filePathArg)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunPy.runScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunPy.runScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunPy_ApplyScript(t *testing.T) {
	type fields struct {
		runtime     string
		scriptsPath string
		storagePath string
	}
	type args struct {
		ctx            context.Context
		list           []string
		script2execute string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ScriptResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &RunPy{
				runtime:     tt.fields.runtime,
				scriptsPath: tt.fields.scriptsPath,
				storagePath: tt.fields.storagePath,
			}
			got, err := me.ApplyScript(tt.args.ctx, tt.args.list, tt.args.script2execute)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunPy.ApplyScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunPy.ApplyScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunPy_ExecuteOne(t *testing.T) {
	type fields struct {
		runtime     string
		scriptsPath string
		storagePath string
	}
	type args struct {
		ctx            context.Context
		script2execute string
		input          string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ScriptResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &RunPy{
				runtime:     tt.fields.runtime,
				scriptsPath: tt.fields.scriptsPath,
				storagePath: tt.fields.storagePath,
			}
			got, err := me.ExecuteOne(tt.args.ctx, tt.args.script2execute, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunPy.ExecuteOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunPy.ExecuteOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
