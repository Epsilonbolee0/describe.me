package domain

import "gorm.io/gorm"

type User struct {
	Login    string `json:"login" gorm:"primaryKey"`
	Password string `json:"password,omitempty"`
	Group    string `json:"group,omitempty"`
	Sex      bool   `json:"sex,omitempty"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`

	Likes              []Function `json:",omitempty" gorm:"many2many:likes"`
	Dislikes           []Function `json:",omitempty" gorm:"many2many:dislikes"`
	PreferredLanguages []Language `json:",omitempty" gorm:"many2many:preferred_languages"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(u).Count(&count)

	if count == 0 {
		tx.Model(u).Update("is_admin", true)
	}

	return
}

type UserAuthDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type UserProfileDTO struct {
	Group string `json:"group"`
	Sex   bool   `json:"sex"`
}
