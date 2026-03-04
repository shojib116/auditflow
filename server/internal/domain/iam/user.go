package iam

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        Email
	FullName     string
	PasswordHash PasswordHash
}
