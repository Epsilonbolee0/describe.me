package domain

type Function struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"code"`
	LanguageID  uint   `json:"language"`
	IsDisplayed bool   `json:"is_displayed" gorm:"default:true"`

	Language Language `json:",omitempty" gorm:"foreignKey:language_id"`
	Likes    []User   `json:",omitempty" gorm:"many2many:likes_on_functions"`
	Dislikes []User   `json:",omitempty" gorm:"many2many:dislikes_on_functions"`
}

type FunctionDTO struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	LanguageID uint   `json:"language"`
}
