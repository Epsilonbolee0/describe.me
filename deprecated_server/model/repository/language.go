package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	Conn *gorm.DB
}

func NewLanguageRepository(conn *gorm.DB) *LanguageRepository {
	return &LanguageRepository{conn}
}

func (repo *LanguageRepository) List() ([]string, error) {
	var languages []string
	err := repo.Conn.Model(&domain.Language{}).Pluck("name", &languages).Error
	return languages, err
}

func (repo *LanguageRepository) Find(name string) (domain.Language, error) {
	var language domain.Language
	err := repo.Conn.Where("name = ?", name).First(&language).Error
	return language, err
}

func (repo *LanguageRepository) Create(language *domain.Language) error {
	return repo.Conn.Create(&language).Error
}

func (repo *LanguageRepository) Delete(id uint) error {
	return repo.Conn.Delete(&domain.Language{ID: id}).Error
}
