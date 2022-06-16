package main

import (
	"errors"

	"github.com/heartBrokenGod/viper-reference/api/handler"
	"github.com/heartBrokenGod/viper-reference/api/server"
	"github.com/heartBrokenGod/viper-reference/repo"
	"github.com/heartBrokenGod/viper-reference/service"
)

func initApiServer(config *Config) (server.ApiServer, error) {
	if config == nil {
		return nil, errors.New("config depencency is nil")
	}
	repo, err := repo.New(config.MySQL)
	if err != nil {
		return nil, err
	}
	service, err := service.New(config.Service, repo)
	if err != nil {
		return nil, err
	}
	handler, err := handler.New(config.Handler, service)
	if err != nil {
		return nil, err
	}
	server, err := server.New(config.Server, handler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
