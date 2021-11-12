package errorz

import (
	"errors"
)

// Common controller errors
var (
	InvalidRequestBody = errors.New("invalid request body")
)
