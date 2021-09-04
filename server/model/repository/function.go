package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type FunctionRepository struct {
	Conn *gorm.DB
}

func NewFunctionRepository(conn *gorm.DB) *FunctionRepository {
	return &FunctionRepository{Conn: conn}
}

func (repo *FunctionRepository) List(languages []uint) ([]domain.Function, error) {
	var functions []domain.Function
	err := repo.Conn.Scopes(isDisplayed, writtenIn(languages)).Order("rating(id) DESC").Find(&functions).Error
	return functions, err
}

func isDisplayed(conn *gorm.DB) *gorm.DB {
	return conn.Where("is_displayed", true)
}

func writtenIn(languages []uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("language_id IN (?)", languages)
	}
}

func (repo *FunctionRepository) Find(id uint) (*domain.Function, error) {
	var function *domain.Function
	err := repo.Conn.First(&function, id).Error
	return function, err
}

func (repo *FunctionRepository) Create(function *domain.Function) error {
	return repo.Conn.Create(&function).Error
}

func (repo *FunctionRepository) UpdateCode(id uint, code string) error {
	return repo.Conn.Model(&domain.Function{ID: id}).Update("code", code).Error
}

func (repo *FunctionRepository) Rating(id uint) int {
	function := domain.Function{ID: id}

	likeCnt := repo.Conn.Model(&function).Association("Likes").Count()
	dislikeCnt := repo.Conn.Model(&function).Association("Dislikes").Count()

	return int(likeCnt - dislikeCnt)
}

func (repo *FunctionRepository) Like(login string, functionID uint) {
	function := domain.Function{ID: functionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&function).Association("Likes").Append(&user)
	repo.Conn.Model(&function).Association("Dislikes").Delete(&user)
}

func (repo *FunctionRepository) Dislike(login string, functionID uint) {
	function := domain.Function{ID: functionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&function).Association("Likes").Delete(&user)
	repo.Conn.Model(&function).Association("Dislikes").Append(&user)
}

func (repo *FunctionRepository) Indifferent(login string, functionID uint) {
	function := domain.Function{ID: functionID}
	user := domain.User{Login: login}

	repo.Conn.Model(&function).Association("Likes").Delete(&user)
	repo.Conn.Model(&function).Association("Dislikes").Delete(&user)
}
