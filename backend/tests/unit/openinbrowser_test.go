package app

import (
	"milonga/milonga/utils"
	"testing"
)

func TestOpenInBrowser(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.OpenInBrowser(tt.args.url)
		})
	}
}
