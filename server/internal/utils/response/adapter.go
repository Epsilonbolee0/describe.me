package response

import (
	. "describe.me/internal/utils/errorz"
	. "describe.me/internal/utils/messagez"
)

import "net/http"

var messageToStatusAdapter = map[string]int{}

const _notFoundStatus = http.StatusBadRequest

func init() {
	bind(http.StatusInternalServerError, InternalError, DatabaseError)
	bind(http.StatusForbidden, LoginAlreadyUsed, UserNotFound, WrongCredentials)
	bind(http.StatusForbidden, LoginIsInvalid, PasswordIsInvalid, EmailIsInvalid)
	bind(http.StatusOK, OK, Created, Found, Updated, Deleted)
	bind(http.StatusNoContent, NotFound)
}

func bind(status int, arguments ...interface{}) {
	for _, arg := range arguments {
		switch value := arg.(type) {
		case error:
			messageToStatusAdapter[value.Error()] = status
		case string:
			messageToStatusAdapter[value] = status
		}
	}
}

func StatusByMessage(msg string) int {
	if status, ok := messageToStatusAdapter[msg]; ok {
		return status
	} else {
		return _notFoundStatus
	}
}
