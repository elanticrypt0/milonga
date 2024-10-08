package utils

import (
	"os/exec"
	"runtime"
)

// open an url in the web browser
func OpenInBrowser(url string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", "explorer", url)
	} else {
		cmd = exec.Command("open", url)
	}

	cmd.Run()
}
