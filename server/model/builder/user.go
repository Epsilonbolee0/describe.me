package builder

import (
	"../domain"
)

type userModifier func(account *domain.User)
type UserBuilder struct {
	actions []userModifier
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) Build() domain.User {
	user := domain.User{}
	for _, action := range b.actions {
		action(&user)
	}

	return user
}

func (b *UserBuilder) Name(name string) *UserBuilder {
	b.actions = append(b.actions, func(user *domain.User) {
		user.Name = name
	})
	return b
}

func (b *UserBuilder) Login(login string) *UserBuilder {
	b.actions = append(b.actions, func(user *domain.User) {
		user.Login = login
	})
	return b
}

func (b *UserBuilder) Password(password string) *UserBuilder {
	b.actions = append(b.actions, func(user *domain.User) {
		user.Password = password
	})
	return b
}

func (b *UserBuilder) Sex(sex bool) *UserBuilder {
	b.actions = append(b.actions, func(user *domain.User) {
		user.Sex = sex
	})
	return b
}

func (b *UserBuilder) IsAdmin() *UserBuilder {
	b.actions = append(b.actions, func(account *domain.User) {
		account.IsAdmin = true
	})
	return b
}
