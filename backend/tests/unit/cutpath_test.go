package app

import (
	"milonga/milonga/utils"
	"testing"
)

func TestCutPath(t *testing.T) {
	type args struct {
		fullPath     string
		subdirectory string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CutPath(tt.args.fullPath, tt.args.subdirectory); got != tt.want {
				t.Errorf("CutPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
