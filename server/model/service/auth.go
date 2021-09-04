package service

import (
	"../../keygen"
	"../../utils"

	"../domain"
	repo "../repository"

	"../builder"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repo.UserRepository
}

func NewAuthService(userRepo *repo.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (auth *AuthService) Login(login, password string) map[string]interface{} {
	user, err := auth.userRepo.Find(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message("User not found")
	default:
		return utils.Message("Error occured while creating user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message("Invalid login credentials")
	}

	user.Password = ""

	resp := utils.Message("Logged In")
	resp["user"] = user

	return resp
}

func (auth *AuthService) Register(key, login, password string) map[string]interface{} {
	if !keygen.IsValid(key) {
		return utils.Message("Key is invalid!")
	}

	userBuilder := builder.NewUserBuilder()
	user := userBuilder.
		Login(login).
		Password(password).
		Build()

	if resp, ok := auth.validate(user); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := auth.userRepo.Create(&user); err != nil {
		return utils.Message("Failure occured while creating user")
	}

	user.Password = ""

	response := utils.Message("User has been created")
	response["user"] = user
	return response
}

func (auth *AuthService) validate(user domain.User) (map[string]interface{}, bool) {
	if !utils.ValidateLogin(user.Login) {
		return utils.Message("Login is invalid"), false
	}

	if !utils.ValidatePassword(user.Password) {
		return utils.Message("Password is invalid"), false
	}

	if !auth.loginIsUnique(user.Login) {
		return utils.Message("Login is already taken"), false
	}

	return utils.Message("Validation passed"), true
}

func (auth *AuthService) loginIsUnique(login string) bool {
	_, err := auth.userRepo.Find(login)
	return err == gorm.ErrRecordNotFound
}
