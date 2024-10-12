package milonga_errors

import (
	"errors"
	"fmt"
	"log"
)

const colorRed = "\033[31m"
const colorReset = "\033[0m"
const errorPrefix = "DBMAN [error] > "

type MilongaError struct {
	StatusCode string
	Err        error
}

func (me *MilongaError) Error() string {
	return fmt.Sprintf("status %s: err %v", me.StatusCode, me.Err)
}

func Generic(code, msg string) *MilongaError {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(msg),
	}
}

func Connection(code, conn string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(GetConnectionError(conn)),
	}
}

func ConnectionFails(code string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(GetFailedToConnect()),
	}
}

func Trying2Connect(code, connname, engine, host, port, user, password, dbname string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(GetTrying2ConnectError(connname, engine, host, port, user, password, dbname)),
	}
}

func Trying2ConnectSQLite(code, connname, engine, dbname string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(GetTrying2ConnectSQLiteError(connname, engine, dbname)),
	}
}

func Trying2ConnectSQLiteFileNotExists(code, filepath string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(FileNotExistError(filepath)),
	}
}

func Instance(code, name, name_lower string) error {
	return &MilongaError{
		StatusCode: code,
		Err:        errors.New(GetInstanceError(name, name_lower)),
	}
}

func FatalErr(msg error) {
	log.Fatalf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, msg)
}

func PrintStr(msg string) {
	log.Printf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, msg)
}

func Print(errorMsg error) {
	log.Printf("%s %s %s %s\n", colorRed, errorPrefix, colorReset, errorMsg)
}
