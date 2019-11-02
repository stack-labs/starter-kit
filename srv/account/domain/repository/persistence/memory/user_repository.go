package memory

import (
	"sync"

	"github.com/micro-in-cn/starter-kit/srv/account/domain/model"
)

type userRepository struct {
	mu    *sync.Mutex
	users map[string]*User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		mu:    &sync.Mutex{},
		users: map[string]*User{},
	}
}

func (r *userRepository) FindByName(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Name == email {
			return model.NewUser(user.ID, user.Name), nil
		}
	}
	return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.GetID()] = &User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}
	return nil
}

type User struct {
	ID   string
	Name string
}
