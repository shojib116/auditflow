package iam

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(strings.ToLower(value))

	_, err := mail.ParseAddress(value)
	if err != nil {
		return "", errors.New("invalid email format")
	}

	return Email(value), nil
}
