package app

import "testing"

func TestGetAppRootPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := path.GetAppRootPath(); got != tt.want {
				t.Errorf("GetAppRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
