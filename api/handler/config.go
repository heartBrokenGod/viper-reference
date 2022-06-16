package handler

type Config struct {
	ValidateRequest bool `yaml:"validate_request"` // dont put required tag on boolean values
}

func NewDefaultConfig() *Config {
	return &Config{
		ValidateRequest: false,
	}
}
