package cli

import (
	"fmt"
	"strings"
)

func PrintBanner(appName string, version string) {
	banner := `
    ╔══════════════════════════════════════╗
    ║     %s                   
    ║     Version: %s                      
    ╚══════════════════════════════════════╝
    `
	// Centrar el nombre de la aplicación
	paddedName := centerText(appName, 30)
	paddedVersion := centerText("v"+version, 30)

	fmt.Printf(banner, paddedName, paddedVersion)
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text
}
