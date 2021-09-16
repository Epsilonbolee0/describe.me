package domain

import (
	"../../utils"
	"gorm.io/gorm"
)

type Description struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Content     string `json:"content" gorm:"not null"`
	AuthorLogin string `json:"author_login"`
	FunctionID  uint   `json:"function_id"`
	IsDisplayed bool   `json:"-" gorm:"default:true"`

	Function Function `json:"-" gorm:"foreignKey:function_id"`
	Author   User     `json:",omitempty" gorm:"foreignKey:author_login"`

	Likes    []User `json:",omitempty" gorm:"many2many:likes_on_descriptions"`
	Dislikes []User `json:",omitempty" gorm:"many2many:dislikes_on_descriptions"`
}

func (d *Description) AfterUpdate(tx *gorm.DB) (err error) {
	var rating int64
	tx.Raw("SELECT * FROM description_rating(?)", d.ID).Scan(&rating)

	if rating >= utils.RATING_THRESHOLD {
		tx.Model(d).Where("id = ?", d.ID).Update("is_displayed", false)
	}

	return
}

type DescribeDTO struct {
	ID      uint   `json:"id"`
	Login   string `json:"login"`
	Content string `json:"string"`
}
