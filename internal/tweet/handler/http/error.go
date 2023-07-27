package http

import (
	"errors"

	"github.com/BasalamahZ/twitter-app/internal/tweet"
)

// Followings are the known errors from User HTTP handlers.
var (
	// errBadRequest is returned when the given request is
	// bad/invalid.
	errBadRequest = errors.New("BAD_REQUEST")

	// errDataNotFound is returned when the desired data is
	// not found.
	errDataNotFound = errors.New("DATA_NOT_FOUND")

	// errInternalServer is returned when there is an
	// unexpected error encountered when processing a request.
	errInternalServer = errors.New("INTERNAL_SERVER_ERROR")

	// errMethodNotAllowed is returned when accessing not
	// allowed HTTP method.
	errMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// errRequestTimeout is returned when processing time
	// has reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")

	// errSourceNotProvided is returned when there is no
	// source provided in the request.
	errSourceNotProvided = errors.New("SOURCE_NOT_PROVIDED")

	// errTooManyRequest is returned when the given request
	// is exceeding the maximum allowed.
	errTooManyRequest = errors.New("TOO_MANY_REQUEST")
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.
	//
	// Internal server error-related should not be mapped
	// here, and the handler should just return `errInternal`
	// as the error instead
	mapHTTPError = map[error]error{
		tweet.ErrDataNotFound:    errDataNotFound,
	}
)
