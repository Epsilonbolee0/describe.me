package service

import (
	"context"
	"regexp"
	"strings"

	"describe.me/internal/objects/entity"
	ts "describe.me/internal/objects/transport"
	"describe.me/internal/service/interfaces"
	"describe.me/internal/service/repository"
	"describe.me/internal/utils/errorz"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService -.
type UserService struct {
	Repo interfaces.UserRepository
}

// New -.
func New(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

// Login - login as either student or teacher
func (serv *UserService) Login(ctx context.Context, dto *ts.UserLoginRequest) (*entity.User, error) {
	user, err := serv.Repo.Find(ctx, dto.Login)
	if err == gorm.ErrRecordNotFound {
		return nil, errorz.UserNotFound
	} else if err != nil {
		return nil, errorz.DatabaseError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errorz.WrongCredentials
	} else if err != nil {
		return nil, errorz.InternalError
	}

	user.Password = ""
	return user, nil
}

// Register - register new student
func (serv *UserService) Register(ctx context.Context, dto *ts.UserRegisterRequest) error {
	user := &entity.User{
		Login:    dto.Login,
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := serv.validate(ctx, user); err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := serv.Repo.Create(context.Background(), user); err != nil {
		return errorz.DatabaseError
	}

	user.Password = ""
	return nil
}

func (serv *UserService) validate(ctx context.Context, user *entity.User) error {
	var err error = nil
	switch {
	case !serv.validateLogin(ctx, user.Login):
		err = errorz.LoginIsInvalid
		fallthrough
	case !serv.validateEmail(user.Email):
		err = errorz.EmailIsInvalid
		fallthrough
	case !serv.validatePassword(user.Password):
		err = errorz.PasswordIsInvalid
	}

	return err
}

func (serv *UserService) validateLogin(ctx context.Context, login string) bool {
	_, err := serv.Repo.Find(ctx, login)
	return err == gorm.ErrRecordNotFound
}

func (serv *UserService) validateEmail(email string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9-_\.]{1,20}$`, email)
	return matched
}

func (serv *UserService) validatePassword(password string) bool {
	capitalLetters := "ABCDEFGHIJKLMNOPQRSUVWXYZ"
	specialSymbols := "-_$.()#!&?/"
	digits := "1234567890"

	return len(password) >= 8 &&
		strings.ContainsAny(password, capitalLetters) &&
		strings.ContainsAny(password, specialSymbols) &&
		strings.ContainsAny(password, digits)
}
