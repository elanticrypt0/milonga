package dbman

import (
	"os"

	"milonga/milonga/dbman/errors"

	"github.com/BurntSushi/toml"
)

func LoadTomlFile[T any](file string, stru *T) {
	if ExitsFile(file) {
		tomlData := string(OpenFile(file))
		_, err := toml.Decode(tomlData, &stru)
		if err != nil {
			errors.FatalErr(err)
		}
	} else {
		errors.PrintStr(errors.FileNotExistError(file))
	}
}

func ExitsFile(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		return false
	}
	return true
}

func OpenFile(file string) []byte {
	if ExitsFile(file) {
		filedata, err := os.ReadFile(file)
		if err != nil {
			errors.PrintStr(errors.FileNotOpened(file))
		}
		return filedata
	} else {
		errors.PrintStr(errors.FileNotExistError(file))
		return nil
	}
}

// RemoveFile elimina un archivo si existe
func RemoveFile(filepath string) error {
	if ExitsFile(filepath) {
		return os.Remove(filepath)
	}
	return nil
}
