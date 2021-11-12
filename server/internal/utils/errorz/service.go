package errorz

import (
	"errors"
)

// Common service errors
var (
	InternalError = errors.New("some service things fucked up")
	DatabaseError = errors.New("error occurred while working with database")
)

// UserService
var (
	LoginAlreadyUsed = errors.New("login already used")

	UserNotFound     = errors.New("user not found")
	WrongCredentials = errors.New("wrong credentials")

	LoginIsInvalid    = errors.New("login is invalid")
	EmailIsInvalid    = errors.New("email is invalid")
	PasswordIsInvalid = errors.New("password is invalid")
)
