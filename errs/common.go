package errs

import "errors"

var (
	// HTTP errors
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")
	DidNotReceiveCSRF = errors.New("did not receive x-csrf-token")

	// Group errors
	ErrRankNotFound     = errors.New("role is over 255 or does not exist")
	ErrRequiresCookie   = errors.New("this endpoint requires a valid cookie")
	ErrGroupDoesntExist = errors.New("group doesn't exist")
	ErrNonOkStatus      = errors.New("response status was not 200") // Only use for endpoints that normally return 200
)
