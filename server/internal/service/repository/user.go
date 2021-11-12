package repository

import (
	"context"
	"describe.me/internal/objects/entity"
	"describe.me/pkg/postgres"
)

type UserRepository struct {
	*postgres.Postgres
}

// NewUserRepository -.
func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (repo *UserRepository) Create(ctx context.Context, user *entity.User) error {
	if err := repo.Conn.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return repo.Auth.AssignRole(user.ID, postgres.Student)
}

func (repo *UserRepository) Find(ctx context.Context, login string) (*entity.User, error) {
	var user *entity.User
	err := repo.Conn.WithContext(ctx).Where("login = ?", login).First(user).Error
	return user, err
}

func (repo *UserRepository) IsTeacher(ctx context.Context, login string) (bool, error) {
	user, err := repo.Find(ctx, login)
	if err != nil {
		return false, err
	}

	isTeacher, err := repo.Auth.CheckRole(user.ID, postgres.Teacher)
	if err != nil {
		return false, err
	}

	return isTeacher, nil
}
