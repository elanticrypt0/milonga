package app

type Config struct {
	Name                   string `toml:"APP_NAME" json:"APP_NAME"`
	Version                string `toml:"APP_VERSION" json:"APP_VERSION"`
	Enviroment             string `toml:"APP_ENVIROMENT" json:"APP_ENVIROMENT"`
	Port                   string `toml:"APP_PORT" json:"APP_PORT"`
	URL                    string `toml:"APP_URL" json:"APP_URL"`
	AppHost                string `toml:"APP_HOST" json:"APP_HOST"`
	APIPath                string `toml:"APP_API_PATH" json:"APP_API_PATH"`
	TemplatesPath          string `toml:"APP_TEMPLATES_PATH" json:"APP_TEMPLATES_PATH"`
	AllowedCORS            string `toml:"APP_ALLOWED_CORS" json:"APP_ALLOWED_CORS"`
	ViewsPath              string `toml:"APP_VIEW_PATH" json:"APP_VIEW_PATH"`
	LogPath                string `toml:"APP_LOG_PATH" json:"APP_LOG_PATH"`
	OpenInBrowser          bool   `toml:"APP_OPENINBROWSER" json:"APP_OPENINBROWSER"`
	DBConfigPath           string `toml:"DB_CONFIG_PATH" json:"DB_CONFIG_PATH"`
	JWTSecret              string `toml:"JWT_SECRET" json:"JWT_SECRET"`
	PasstokenEncryptionkey string `toml:"PASSTOKEN_ENCRYPTIONKEY" json:"PASSTOKEN_ENCRYPTIONKEY"`
}
