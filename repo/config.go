package repo

type Config struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

func NewDefaultConfig() *Config {
	return nil // return nil if all config vars are mandatory
}
