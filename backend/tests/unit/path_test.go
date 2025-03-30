package app

import (
	"milonga/milonga/utils"
	"os"
	"strings"
	"testing"
)

func TestGetAppRootPath(t *testing.T) {
	// Obtener el directorio actual usando el paquete os
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %v", err)
	}
	
	tests := []struct {
		name string
		want string
	}{
		{
			name: "returns current working directory",
			want: expected,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GetAppRootPath()
			// Asegúrate de que las rutas estén normalizadas antes de comparar
			if !strings.EqualFold(got, tt.want) {
				t.Errorf("GetAppRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
