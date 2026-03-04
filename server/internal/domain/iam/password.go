package iam

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordHash string

func NewPasswordHash(plain string) (PasswordHash, error) {
	if !isPasswordValid(plain) {
		return "", ErrInvalidPassword
	}

	hash, err := HashPassword(plain)
	if err != nil {
		return "", err
	}

	return PasswordHash(hash), nil
}

var passwordRegex = regexp.MustCompile(
	`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/\?]).{8,}$`,
)

func isPasswordValid(plain string) bool {
	return passwordRegex.MatchString(plain)
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
