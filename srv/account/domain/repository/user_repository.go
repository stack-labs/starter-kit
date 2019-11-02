package repository

import "github.com/micro-in-cn/starter-kit/srv/account/domain/model"

type UserRepository interface {
	FindByName(name string) (*model.User, error)
	Save(*model.User) error
}
