package errs

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest = errors.New("bad request")
)
