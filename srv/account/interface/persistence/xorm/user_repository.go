package xorm

import (
	"github.com/micro-in-cn/starter-kit/srv/account/domain/model"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(id int64) (*model.User, error) {
	user := model.User{}
	if ok, err := DB().Where("id = ?", id).Get(&user); ok && err == nil {
		return &user, nil
	} else {
		return nil, err
	}
}

func (r *userRepository) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if has, err := DB().Where("name = ?", name).Get(&user); err == nil && has {
		return &user, nil
	} else {
		return nil, err
	}
}

func (r *userRepository) Add(user *model.User) error {
	id, err := DB().Insert(user)
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}

func (r *userRepository) List(page, size int) ([]*model.User, error) {
	list := make([]*model.User, 0)
	session := DB().Desc("id")
	err := session.Limit(size, (page-1)*size).Find(&list)

	return list, err
}
