package model

type User struct {
	id       string
	name     string
	password string
}

func NewUser(id, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetPassword() string {
	return u.password
}
