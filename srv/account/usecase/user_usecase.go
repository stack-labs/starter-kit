package usecase

import (
	"github.com/google/uuid"
	"github.com/micro-in-cn/starter-kit/srv/account/domain/model"
	"github.com/micro-in-cn/starter-kit/srv/account/domain/repository"
	"github.com/micro-in-cn/starter-kit/srv/account/domain/service"
)

type UserUsecase interface {
	LoginUser(name, pwd string) (*User, error)
	RegisterUser(email string) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) LoginUser(name, pwd string) (*User, error) {
	user, err := u.repo.FindByName(name)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	return &User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}, nil
}

func (u *userUsecase) RegisterUser(name string) error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	if err := u.service.Duplicated(name); err != nil {
		return err
	}
	user := model.NewUser(uid.String(), name)
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

type User struct {
	ID   string
	Name string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:   user.GetID(),
			Name: user.GetName(),
		}
	}
	return res
}
