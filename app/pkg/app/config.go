package app

type Config struct {
	Name         string `toml:"APP_NAME"`
	Version      string `toml:"APP_VERSION"`
	Port         string `toml:"APP_PORT"`
	URL          string `toml:"APP_URL"`
	AppHost      string
	ViewsPath    string `toml:"APP_VIEW_PATH"`
	LogPath      string `toml:"APP_LOG_PATH"`
	DBConfigPath string `toml:"DB_CONFIG_PATH"`
}
