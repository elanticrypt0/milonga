package app

type Config struct {
	Name    string `toml:"APP_NAME"`
	Version string `toml:"APP_VERSION"`
	Port    string `toml:"APP_PORT"`
}
