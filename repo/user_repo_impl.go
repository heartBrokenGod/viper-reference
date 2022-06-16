package repo

import (
	"errors"
	"fmt"

	"github.com/heartBrokenGod/viper-reference/entity"
)

type UserRepoImpl struct {
	config         *Config
	data           map[uint32]*entity.User
	lastInsertedID uint32
}

func New(config *Config) (*UserRepoImpl, error) {
	if config == nil {
		return nil, errors.New("config dependency is nil")
	}
	if config.User != "root" || config.Password != "password" {
		return nil, errors.New("database connection err: invalid credentials")
	}
	return &UserRepoImpl{
		config:         config,
		data:           map[uint32]*entity.User{},
		lastInsertedID: 0,
	}, nil
}

func (userRepoImpl *UserRepoImpl) Read(id uint32) (*entity.User, error) {
	user, found := userRepoImpl.data[id]
	if !found {
		return nil, errors.New(fmt.Sprint("no data exist for user with id: ", id))
	}
	return user, nil
}

func (userRepoImpl *UserRepoImpl) Create(user entity.User) (*entity.User, error) {
	userRepoImpl.lastInsertedID++
	user.ID = userRepoImpl.lastInsertedID
	userRepoImpl.data[user.ID] = &user
	return &user, nil
}
