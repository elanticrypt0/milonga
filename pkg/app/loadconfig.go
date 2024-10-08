package app

import (
	"milonga/pkg/errors"
	"milonga/pkg/utils"

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
			errors.FatalErr(err)
		}
	} else {
		errors.PrintStr(errors.FileNotExistError(file))
	}
}