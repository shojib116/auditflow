package iam

import (
	"context"
	"errors"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email Email) (*User, error)
}

var (
	ErrUserAlreadyExists = errors.New("an user with this email already exists!")
	ErrInvalidPassword   = errors.New("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
	ErrUserNotFound      = errors.New("no user can be found with this email!")
)
