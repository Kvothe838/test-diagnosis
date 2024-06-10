package graceful

import "errors"

var (
	ErrAlreadyRegistered   = errors.New("service already registered")
	ErrTimeout             = errors.New("graceful shutdown timeout")
	ErrRegistrationsClosed = errors.New("registrations are closed")
	ErrNotComparable       = errors.New("type is not comparable")
)
