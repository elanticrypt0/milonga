package utils

import (
	"log"
	"os"
)

func GetAppRootPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return pwd

}
