package app

type DefaultAdmin struct {
	Email         string `toml:"DEFAULT_USER_ADMIN_EMAIL"`
	Password      string `toml:"DEFAULT_USER_ADMIN_PASSW"`
	Username      string `toml:"DEFAULT_USER_ADMIN_USERNAME"`
}
