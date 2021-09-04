package domain

type Language struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`

	PreferredLanguages []User `json:",omitempty" gorm:"many2many:preferred_languages"`
}

type LanguageDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserPreferrefLanguagesDTO struct {
	ID uint `json:"id"`
}
