package milonga_errors

import (
	"fmt"
)

func FileNotExistError(filepath string) string {
	return fmt.Sprintf("File doesn't exists: %q \n", filepath)
}

func FileNotOpened(filepath string) string {
	return fmt.Sprintf("Cann't open file: %q \n", filepath)
}

func FileNotLoaded(filepath string) string {
	return fmt.Sprintf("Cann't load or parse file: %q \n", filepath)
}

func GetConnectionError(conn string) string {
	return fmt.Sprintf("The connection %s does not exists \n", conn)
}

func GetFailedToConnect() string {
	return "DB connection fails."
}

func GetTrying2ConnectError(connname, engine, host, port, user, password, dbname string) string {
	return fmt.Sprintf("Name: %q, Engine: %q host: %q, port: %q, user: %q, password: %q, database name: %q \n", connname, engine, host, port, user, password, dbname)
}

func GetTrying2ConnectSQLiteError(connname, engine, dbname string) string {
	return fmt.Sprintf("Name: %q, Engine: %q database path: %q \n", connname, engine, dbname)
}

func GetInstanceError(name, name_lower string) string {
	return fmt.Sprintf("The connection %s (%s) does not exists \n", name, name_lower)
}
