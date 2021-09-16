package domain

type Description struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Content     string `json:"content"`
	AuthorLogin string `json:"author_login"`
	FunctionID  uint   `json:"function_id"`
	IsDisplayed bool   `json:"is_displayed" gorm:"default:true"`

	Function Function `json:",omitempty" gorm:"foreignKey:function_id"`
	Author   User     `json:",omitempty" gorm:"foreignKey:author_login"`

	Likes    []User `json:",omitempty" gorm:"many2many:likes_on_descriptions"`
	Dislikes []User `json:",omitempty" gorm:"many2many:dislikes_on_descriptions not null"`
}

type DescribeDTO struct {
	ID      uint   `json:"id"`
	Login   string `json:"login"`
	Content string `json:"string"`
}
