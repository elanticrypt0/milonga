package db

import (
	"fmt"
	"log"
	"milonga/utils"

	"github.com/elanticrypt0/dbman"
)

const dbConfigFile = "db_config.toml"

func ConnectDB(connectionName string) *dbman.DBMan {
	// conexion a la db
	db := dbman.New()
	dbConfigPath := utils.GetAppRootPath() + "config/"
	fmt.Printf("DB config path: %q\n", dbConfigPath)
	db.SetRootPath(dbConfigPath)
	db.LoadConfigToml(dbConfigPath + dbConfigFile)
	err := db.Connect("local")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	db.SetPrimary(connectionName)

	db.Primary.AutoMigrate()

	return db
}
