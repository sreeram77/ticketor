package errors

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalid      = errors.New("invalid")
	ErrNotAvailable = errors.New("not available")
)
