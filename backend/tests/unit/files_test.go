package app

import (
	"milonga/milonga/utils"
	"reflect"
	"testing"
)

func TestLoadTomlFile(t *testing.T) {
	type args struct {
		file string
		stru *T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.LoadTomlFile(tt.args.file, tt.args.stru)
		})
	}
}

func TestLoadJSONFile(t *testing.T) {
	type args struct {
		file string
		stru *T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.LoadJSONFile(tt.args.file, tt.args.stru)
		})
	}
}

func TestExitsFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ExitsFile(tt.args.filepath); got != tt.want {
				t.Errorf("ExitsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.OpenFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
