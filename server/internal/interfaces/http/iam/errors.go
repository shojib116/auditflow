package iam

import (
	"errors"
	"fmt"
	"net/http"

	iamDomain "github.com/shojib116/auditflow-api/internal/domain/iam"
)

type AppError struct {
	Err        error
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func MapError(err error) *AppError {
	switch {
	case errors.Is(err, iamDomain.ErrUserAlreadyExists):
		return &AppError{
			Err:        err,
			Message:    fmt.Sprintf("User already exists: %s", err.Error()),
			StatusCode: http.StatusConflict,
		}
	case errors.Is(err, iamDomain.ErrInvalidPassword):
		return &AppError{
			Err:        err,
			Message:    fmt.Sprintf("password is invalid: %s", err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	case errors.Is(err, iamDomain.ErrUserNotFound):
		return &AppError{
			Err:        err,
			Message:    fmt.Sprintf("user not found: %s", err.Error()),
			StatusCode: http.StatusNotFound,
		}
	default:
		return &AppError{
			Err:        err,
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}
	}
}
