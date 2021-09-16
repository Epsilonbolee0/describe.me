package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type DescriptionRepository struct {
	Conn *gorm.DB
}

func NewDescriptionRepository(conn *gorm.DB) *DescriptionRepository {
	return &DescriptionRepository{conn}
}

func (repo *DescriptionRepository) ListByFunction(functionID uint) ([]string, error) {
	var descriptions []string
	err := repo.Conn.Model(&domain.Description{}).Where("function_id = ?", functionID).Pluck("content", &descriptions).Error
	return descriptions, err
}

func (repo *DescriptionRepository) Create(description *domain.Description) error {
	return repo.Conn.Create(&description).Error
}

func (repo *DescriptionRepository) Delete(id uint) error {
	return repo.Conn.Delete(&domain.Description{ID: id}).Error
}

func (repo *DescriptionRepository) Rating(id uint) int {
	description := domain.Description{ID: id}

	likeCnt := repo.Conn.Model(&description).Association("Likes").Count()
	dislikeCnt := repo.Conn.Model(&description).Association("Dislikes").Count()

	return int(likeCnt - dislikeCnt)
}

func (repo *DescriptionRepository) Like(login string, descriptionID uint) {
	description := domain.Description{ID: descriptionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&description).Association("Likes").Append(&user)
	repo.Conn.Model(&description).Association("Dislikes").Delete(&user)
}

func (repo *DescriptionRepository) Dislike(login string, descriptionID uint) {
	description := domain.Description{ID: descriptionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&description).Association("Likes").Delete(&user)
	repo.Conn.Model(&description).Association("Dislikes").Append(&user)
}

func (repo *DescriptionRepository) Indifferent(login string, descriptionID uint) {
	description := domain.Description{ID: descriptionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&description).Association("Likes").Delete(&user)
	repo.Conn.Model(&description).Association("Dislikes").Delete(&user)
}
