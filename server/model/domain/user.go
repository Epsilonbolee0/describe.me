package domain

type User struct {
	Login    string `json:"login" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Group    string `json:"group"`
	Sex      bool   `json:"sex"`
	IsAdmin  bool   `json:"is_moderator" gorm:"default:false"`
}

type UserAuthDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type UserProfileDTO struct {
	Name  string `json:"name"`
	Group string `json:"group"`
	Sex   bool   `json:"sex"`
}
