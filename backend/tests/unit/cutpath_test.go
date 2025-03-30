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
		{
			name: "standard path cutting",
			args: args{
				fullPath:     "/home/user/projects/milonga/backend",
				subdirectory: "milonga",
			},
			want: "milonga/backend",
		},
		{
			name: "subdirectory at the beginning",
			args: args{
				fullPath:     "milonga/backend/src",
				subdirectory: "milonga",
			},
			want: "milonga/backend/src",
		},
		{
			name: "subdirectory not found",
			args: args{
				fullPath:     "/home/user/projects/other/backend",
				subdirectory: "milonga",
			},
			want: "/home/user/projects/other/backend",
		},
		{
			name: "empty subdirectory",
			args: args{
				fullPath:     "/home/user/projects/milonga/backend",
				subdirectory: "",
			},
			want: "/home/user/projects/milonga/backend",
		},
		{
			name: "empty full path",
			args: args{
				fullPath:     "",
				subdirectory: "milonga",
			},
			want: "",
		},
		{
			name: "multiple occurrences of subdirectory",
			args: args{
				fullPath:     "/home/milonga/projects/milonga/backend",
				subdirectory: "milonga",
			},
			want: "milonga/projects/milonga/backend",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CutPath(tt.args.fullPath, tt.args.subdirectory); got != tt.want {
				t.Errorf("CutPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
