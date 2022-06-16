package service

import "github.com/heartBrokenGod/viper-reference/entity"

type UserRepo interface {
	Read(id uint32) (*entity.User, error)
	Create(entity.User) (*entity.User, error)
}
