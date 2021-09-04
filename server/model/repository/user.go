package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{Conn: conn}
}

func (repo *UserRepository) Create(user *domain.User) error {
	return repo.Conn.Create(&user).Error
}

func (repo *UserRepository) Find(login string) (domain.User, error) {
	var user domain.User
	err := repo.Conn.Where("login = ?", login).First(&user).Error
	return user, err
}

func (repo *UserRepository) UpdateName(login string, name string) error {
	return repo.Conn.Model(&domain.User{Login: login}).Update("name", name).Error
}

func (repo *UserRepository) UpdateGroup(login string, group string) error {
	return repo.Conn.Model(&domain.User{Login: login}).Update("group", group).Error
}

func (repo *UserRepository) UpdateSex(login string, sex bool) error {
	return repo.Conn.Model(&domain.User{Login: login}).Update("sex", sex).Error
}
