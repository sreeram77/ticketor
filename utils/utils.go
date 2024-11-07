package utils

import (
	"errors"
	"net/mail"

	intErrors "ticketor/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckEmail checks if email is valid.
func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}

// StatusFromError returns status from error.
func StatusFromError(err error) error {
	switch {
	case errors.Is(err, intErrors.ErrNotFound):
		return status.Errorf(codes.NotFound, err.Error())
	case errors.Is(err, intErrors.ErrInvalid):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, intErrors.ErrNotAvailable):
		return status.Errorf(codes.Internal, err.Error())
	default:
		return status.Errorf(codes.Unknown, err.Error())
	}
}
