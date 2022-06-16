package server

import "time"

type Config struct {
	TimeOut time.Duration `validate:"required"`
	Port    int
}

func NewDefaultConfig() *Config {
	return &Config{
		TimeOut: time.Minute * 1,
		Port:    8080, // default
	}
}
