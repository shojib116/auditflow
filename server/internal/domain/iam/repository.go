package iam

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email Email) (*User, error)
}
