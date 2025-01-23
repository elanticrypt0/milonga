package dbman

type DBConfigSlice struct {
	Configurations []DBConfig `toml:"database"`
}
