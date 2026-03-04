package iam

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/crypto/argon2"
)

type PasswordHash string

func NewPasswordHash(plain string) (PasswordHash, error) {
	if err := isPasswordValid(plain); err != nil {
		return "", err
	}

	hash, err := HashPassword(plain)
	if err != nil {
		return "", err
	}

	return PasswordHash(hash), nil
}

func isPasswordValid(p string) error {
	if len(p) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, c := range p {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()_+-=[]{};':\"\\|,.<>/?", c):
			hasSpecial = true
		}
	}

	if !hasLower {
		return errors.New("password must contain a lowercase letter")
	}
	if !hasUpper {
		return errors.New("password must contain an uppercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain a digit")
	}
	if !hasSpecial {
		return errors.New("password must contain a special character")
	}

	return nil
}

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var defaultParams = ArgonParams{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, defaultParams.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultParams.Iterations,
		defaultParams.Memory,
		defaultParams.Parallelism,
		defaultParams.KeyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, defaultParams.Memory, defaultParams.Iterations, defaultParams.Parallelism, b64Salt, b64Hash)

	return encoded, nil
}

func ComparePassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var p ArgonParams
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	comparisonHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}
