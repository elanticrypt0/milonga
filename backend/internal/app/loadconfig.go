package app

import (
	"milonga/internal/milonga_errors"
	"milonga/internal/utils"

	"github.com/BurntSushi/toml"
)

func LoadConfig(configPath string) *Config {
	config := &Config{}
	LoadTomlFile(configPath, config)
	return config
}

func LoadTomlFile[T any](file string, stru *T) {
	if utils.ExitsFile(file) {
		tomlData := string(utils.OpenFile(file))
		_, err := toml.Decode(tomlData, &stru)
		if err != nil {
			milonga_errors.FatalErr(err)
		}
	} else {
		milonga_errors.PrintStr(milonga_errors.FileNotExistError(file))
	}
}
