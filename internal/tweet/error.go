package tweet

import "errors"

// Followings are the known errors returned from user.
var (
	// ErrDataNotFound is returned when the wanted data is
	// not found.
	ErrDataNotFound = errors.New("data not found")

	// ErrInvalidStatus is returned when the given study program status
	// is invalid.
	ErrInvalidStatus = errors.New("invalid status")
)
