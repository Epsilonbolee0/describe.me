package domain

import (
	"../../utils"
	"gorm.io/gorm"
)

type Function struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"code" gorm:"not null"`
	LanguageID  uint   `json:"language" gorm:"not null"`
	IsDisplayed bool   `json:"-" gorm:"default:true"`

	Language Language `json:"-" gorm:"foreignKey:language_id"`
	Likes    []User   `json:",omitempty" gorm:"many2many:likes_on_functions"`
	Dislikes []User   `json:",omitempty" gorm:"many2many:dislikes_on_functions"`
}

func (f *Function) AfterUpdate(tx *gorm.DB) (err error) {
	var rating int64
	tx.Raw("SELECT * FROM function_rating(?)", f.ID).Scan(&rating)

	if rating >= utils.RATING_THRESHOLD {
		tx.Model(f).Where("id = ?", f.ID).Update("is_displayed", false)
	}

	return
}

type FunctionDTO struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	LanguageID uint   `json:"language"`
}
