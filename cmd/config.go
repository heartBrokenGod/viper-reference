package main

import (
	"flag"
	"strings"

	"github.com/go-playground/validator"
	"github.com/heartBrokenGod/viper-reference/api/handler"
	"github.com/heartBrokenGod/viper-reference/api/server"
	"github.com/heartBrokenGod/viper-reference/repo"
	"github.com/heartBrokenGod/viper-reference/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Server  *server.Config  `validate:"required"`
	Handler *handler.Config `validate:"required"`
	Service *service.Config `validate:"required"`
	MySQL   *repo.Config    `validate:"required"`
}

func initConfig() (*Config, error) {
	config := &Config{ // Create the instance of the composed configuration struct
		Server:  server.NewDefaultConfig(),
		Handler: handler.NewDefaultConfig(),
		MySQL:   repo.NewDefaultConfig(),
	}
	viper.SetConfigName("config")             // set the config file name
	viper.SetConfigType("yaml")               // set the config file format
	viper.AddConfigPath(".")                  // set the file path where config can be found
	viper.SetEnvPrefix("VIPER")               // set the env var prefix
	replacer := strings.NewReplacer(".", "_") // linux does not support env vars names with .(dot)
	viper.SetEnvKeyReplacer(replacer)         // thats why . is replaced with _
	viper.AutomaticEnv()

	// specify the flag to get values from command line
	flag.Int("server.port", 8080, "application port to listen on")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// read the config
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// unmarshal config data into the config instance
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	// validate the config
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return nil, err
	}
	// viper.WriteConfigAs("./copy_config.json") // for the demo purpose

	return config, nil
}
