package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (repo *UserRepository) Create(user *domain.User) error {
	return repo.Conn.Create(&user).Error
}

func (repo *UserRepository) Find(login string) (domain.User, error) {
	var user domain.User
	err := repo.Conn.Where("login = ?", login).First(&user).Error
	return user, err
}

func (repo *UserRepository) ListPreferredLanguages(login string) []domain.Language {
	var languages []domain.Language
	user := domain.User{Login: login}
	repo.Conn.Model(&user).Association("PreferredLanguages").Find(&languages)
	return languages
}

func (repo *UserRepository) AddPreferredLanguage(login string, id uint) {
	language := domain.Language{ID: id}
	user := domain.User{Login: login}
	repo.Conn.Model(&user).Association("PreferredLanguages").Append(&language)
}

func (repo *UserRepository) DeletePreferredLanguage(login string, id uint) {
	language := domain.Language{ID: id}
	user := domain.User{Login: login}
	repo.Conn.Model(&user).Association("PreferredLanguages").Delete(&language)
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
