package dbman

type DBConfig struct {
	ConnName string `toml:"connName"`
	Engine   string `toml:"engine"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

func NewDBConfig(connname, engine, host, port, user, passwd, dbname string) DBConfig {

	return DBConfig{
		ConnName: connname,
		Engine:   engine,
		Host:     host,
		Port:     port,
		User:     user,
		Password: passwd,
		DBName:   dbname,
	}
}
