package usecase

import (
	"github.com/hb-go/pkg/conv"
	"github.com/micro-in-cn/starter-kit/console/account/domain/repository"
	"github.com/micro-in-cn/starter-kit/console/account/domain/service"
)

type UserUsecase interface {
	LoginUser(name, pwd string) (*User, error)
	RegisterUser(name, pwd string) (*User, error)
	GetUser(id int64) (*User, error)
	GetUserList(page, size int) ([]*User, error)
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

func (this *userUsecase) LoginUser(name, pwd string) (*User, error) {
	user, err := this.service.Login(name, pwd)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	u := &User{}
	conv.StructToStruct(user, u)
	return u, nil
}

func (this *userUsecase) RegisterUser(name, pwd string) (*User, error) {
	user, err := this.service.Register(name, pwd)
	if err != nil {
		return nil, err
	}

	u := &User{}
	conv.StructToStruct(user, u)
	return u, nil
}

func (this *userUsecase) GetUser(id int64) (*User, error) {
	user, err := this.repo.FindById(id)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	}

	return &User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (this *userUsecase) GetUserList(page, size int) ([]*User, error) {
	list, err := this.repo.List(page, size)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, len(list))
	for _, u := range list {
		user := &User{}
		conv.StructToStruct(u, user)
		users = append(users, user)
	}

	return users, nil
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
