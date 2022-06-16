package service

import (
	"errors"

	"github.com/heartBrokenGod/viper-reference/entity"
)

type UserServiceImpl struct {
	config *Config
	repo   UserRepo
}

func New(config *Config, repo UserRepo) (*UserServiceImpl, error) {
	if config == nil {
		return nil, errors.New("config dependency is nil")
	}
	if repo == nil {
		return nil, errors.New("repo dependency is nil")
	}
	return &UserServiceImpl{
		config: config,
		repo:   repo,
	}, nil
}

func (userServiceImpl *UserServiceImpl) Get(id uint32) (*entity.User, error) {
	// since service layer not doing much
	// can directly call the repo layer
	return userServiceImpl.repo.Read(id)
}

func (userServiceImpl *UserServiceImpl) Add(user entity.User) (*entity.User, error) {
	// check for config allow invalid username
	if !userServiceImpl.config.AllowInvalidUserName {
		if len(user.Name) <= 1 {
			return nil, errors.New("could not allow to add invalid user name")
		}
	}
	return userServiceImpl.repo.Create(user)
}
