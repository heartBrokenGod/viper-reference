package service

import "github.com/heartBrokenGod/viper-reference/entity"

type UserService interface {
	Get(id uint32) (*entity.User, error)
	Add(entity.User) (*entity.User, error)
}
