package iam

import (
	"context"
	"errors"

	iamDomain "github.com/shojib116/auditflow-api/internal/domain/iam"
)

type UserService struct {
	repo iamDomain.UserRepository
}

func NewUserService(r iamDomain.UserRepository) UserService {
	return UserService{
		repo: r,
	}
}

type RegisterRequestInput struct {
	Email    string
	Password string
	FullName string
}

func (s *UserService) RegisterUser(ctx context.Context, input RegisterRequestInput) (*iamDomain.User, error) {
	email, err := iamDomain.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.GetUserByEmail(ctx, email)
	if err == nil {
		return nil, iamDomain.ErrUserAlreadyExists
	}

	if !errors.Is(err, iamDomain.ErrUserNotFound) {
		return nil, err
	}

	hash, err := iamDomain.NewPasswordHash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, &iamDomain.User{
		Email:        email,
		PasswordHash: hash,
		FullName:     input.FullName,
	})
	if err != nil {
		if errors.Is(err, iamDomain.ErrUserAlreadyExists) {
			return nil, iamDomain.ErrUserAlreadyExists
		}
		return nil, err
	}

	return user, nil
}
