package utils

import (
	"milonga/pkg/milonga_errors"
	"os"

	"github.com/BurntSushi/toml"
)

func LoadTomlFile[T any](file string, stru *T) {
	if ExitsFile(file) {
		tomlData := string(OpenFile(file))
		_, err := toml.Decode(tomlData, &stru)
		if err != nil {
			milonga_errors.FatalErr(err)
		}
	} else {
		milonga_errors.PrintStr(milonga_errors.FileNotExistError(file))
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
			milonga_errors.PrintStr(milonga_errors.FileNotOpened(file))
		}
		return filedata
	} else {
		milonga_errors.PrintStr(milonga_errors.FileNotExistError(file))
		return nil
	}
}
