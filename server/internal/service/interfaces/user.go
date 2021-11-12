package interfaces

import (
	"context"
	"describe.me/internal/objects/entity"
	ts "describe.me/internal/objects/transport"
)

type (
	//UserService -.
	UserService interface {
		Register(ctx context.Context, dto *ts.UserRegisterRequest) error
		Login(ctx context.Context, dto *ts.UserLoginRequest) (*entity.User, error)
	}

	//UserRepository -.
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) error
		Find(ctx context.Context, login string) (*entity.User, error)
		IsTeacher(ctx context.Context, login string) (bool, error)
	}
)
