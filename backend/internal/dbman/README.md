# DBMan

This a db conection manager now supports: **SQLite, MySQL y PostgresSQL**
This lib the amazing [GORM](https://gorm.io/)

The magik of this package is every time that you need get a database connections instance you can get an error and do stuff if that happen.

This makes posible that your applications doesn't panic if the db connection is wrong.

# Install

```go
go get dbman
```

# Config

You can get the database conexions from .env or .toml file. Personaly I recommend use toml and you can write there several database conexions.

## Config from .toml file

```go
[db.conn1]
    connName="conn1"
    engine="mysql"
    host="localhost"
    port="3336"
    user="root"
    password="secret"
    dbname="homestead"

[db.local]
    connName="local"
    engine="sqlite"
    host=""
    port=""
    user=""
    password=""
    dbname="my_db"
```

## Config from .env file

```go
DB_CONN_NAME = ""
DB_ENGINE = ""
DB_HOST = ""
DB_PORT = ""
DB_USER = ""
DB_PASSWORD = ""
DB_NAME = ""
```

# Use example

## Load config

```go
// instace the DBManager
dbman:=dbman.New()
dbman.LoadConfigToml("./config/db.toml")

// Connects the first database.
// By default is setted as primary
dbman.Connect("Conn1")

// this instance could be called like this
// this retuns an *gorm.DB
dbman.Primary

// or can be called like this
// this also returns an *gorm.DB
dbman.GetInstance("conn1")

// You can set a secondary connection
dbman.Connect("local")
dbman.SetSecondary("local")
dbman.GetInstance("local")

```

# Load config from env file

```go
dbman.LoadConfigEnv("./env")
```

# Connections errors

This is important. Every time that you call GetInstance this checks if the instances has no errors. If the connectio have errors or is not connected returns the error or nil

## The instance is not connected

For example if you use:

```go
primary:= dbman.Primary
```

Primary would be **nil**

## The connection has fail

If the connection has some credential or other error the

```go
db,err:= dbman.GetInstance("conn1")
if err!=nil{
    log.Println(err)
}
// ...
```

## Debug functions

Are three functions to make a better debuggin. These functions print information of the instance:

- PrintConnectionsList
- PrintActiveConnectionsList
- CheckDefaultConnections
