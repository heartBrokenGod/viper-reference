package service

type Config struct {
	AllowInvalidUserName bool
}

func NewDefaultConfig() *Config {
	return &Config{
		AllowInvalidUserName: false,
	}
}
