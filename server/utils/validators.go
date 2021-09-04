package utils

import (
	"regexp"
	"strings"
)

func ValidateLogin(login string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9-_\.]{1,20}$`, login)
	return matched
}

func ValidatePassword(password string) bool {
	capitalLetters := "ABCDEFGHIJKLMNOPQRSUVWXYZ"
	specialSymbols := "-_$.()#!&?/"
	digits := "1234567890"

	return len(password) >= 8 && strings.ContainsAny(password, capitalLetters) && strings.ContainsAny(password, specialSymbols) && strings.ContainsAny(password, digits)
}
