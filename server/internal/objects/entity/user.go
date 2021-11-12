package entity

// User -.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
