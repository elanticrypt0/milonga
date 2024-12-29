package utils

import "strings"

// removes elements of some path until the subdirectory

func CutPath(fullPath, subdirectory string) string {
	index := strings.Index(fullPath, subdirectory)
	if index != -1 {
		return fullPath[index:]
	}
	return fullPath
}
