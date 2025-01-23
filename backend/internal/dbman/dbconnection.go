package dbman

import (
	"fmt"
	"milonga/internal/dbman/console"

	"milonga/internal/dbman/errors"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnection struct {
	DBConfig DBConfig
	Instance *gorm.DB
	ErrConn  error
}

// Create new DB connection
func NewDBConn(connData DBConfig) DBConnection {
	return DBConnection{
		DBConfig: connData,
		Instance: nil,
		ErrConn:  nil,
	}
}

// Check if the connection has errors
func (me *DBConnection) IsOk() bool {
	return me.ErrConn == nil
}

// Checks if the connection is connected
func (me *DBConnection) IsConnected() bool {
	return me.Instance != nil
}

// Connects the database by engine
// Mysql
// Postgres
// SQLite
func (me *DBConnection) Connect(rootPath string) error {

	switch me.DBConfig.Engine {
	case "mysql":
		conn, err := me.connect2Mysql()
		if err != nil {
			return err
		}
		me.Instance = conn
	case "postgres":
		conn, err := me.connect2Postgres()
		if err != nil {
			return err
		}
		me.Instance = conn
	case "sqlite":
		conn, err := me.connect2SQLite(rootPath)
		if err != nil {
			return err
		}
		me.Instance = conn
	default:
		errors.PrintStr(fmt.Sprintf("The connections engine %q is not valid.\n", me.DBConfig.Engine))
		return nil
	}

	return nil
}

// Connects to mysql
func (me *DBConnection) connect2Mysql() (*gorm.DB, error) {
	const dns = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dnsConfig := fmt.Sprintf(dns, me.DBConfig.User, me.DBConfig.Password, me.DBConfig.Host, me.DBConfig.Port, me.DBConfig.DBName)
	// connect to gorn
	conn, err := gorm.Open(mysql.Open(dnsConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		me.ErrConn = me.returnErrConn()
		return nil, me.ErrConn
	}
	return conn, nil
}

// Connects to postgres
func (me *DBConnection) connect2Postgres() (*gorm.DB, error) {

	const dns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable "
	dnsConfig := fmt.Sprintf(dns, me.DBConfig.Host, me.DBConfig.User, me.DBConfig.Password, me.DBConfig.DBName, me.DBConfig.Port)
	// connect to gorn
	conn, err := gorm.Open(postgres.Open(dnsConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		me.ErrConn = me.returnErrConn()
		return nil, me.ErrConn
	}
	// todo
	return conn, nil
}

// connects to sqlite
func (me *DBConnection) connect2SQLite(rootPath string) (*gorm.DB, error) {
	dbname := me.DBConfig.DBName
	dbpath := "./" + dbname
	// if the file doesnt exists then create the file and make the connection
	if !ExitsFile(dbpath) {
		errors.Print(errors.Trying2ConnectSQLiteFileNotExists("66", dbpath))
		// remove the first dot in case is "./_db/local"
		dbpath = rootPath + "/" + dbname
		if !ExitsFile(dbpath) {
			errors.Print(errors.Trying2ConnectSQLiteFileNotExists("67", dbpath))
			console.Print(fmt.Sprintf("Creating new %q", dbpath))
		}
	}
	// gorm create sqlite db
	conn, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		me.ErrConn = errors.Trying2ConnectSQLite("0", me.DBConfig.ConnName, me.DBConfig.Engine, dbpath)
		return nil, me.ErrConn
	}
	return conn, nil
}

// Return the errors connection
func (me *DBConnection) returnErrConn() error {
	errors.FatalErr(errors.Trying2Connect("400", me.DBConfig.ConnName, me.DBConfig.Engine, me.DBConfig.Host, me.DBConfig.Port, me.DBConfig.User, me.DBConfig.Password, me.DBConfig.DBName))
	return errors.ConnectionFails("0")
}
