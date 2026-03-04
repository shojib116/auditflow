package iam

import (
	"context"
	"database/sql"

	"github.com/shojib116/auditflow-api/internal/database"
	iamDomain "github.com/shojib116/auditflow-api/internal/domain/iam"
)

type userRepository struct {
	db      *sql.DB
	queries *database.Queries
}

func NewUserRepository(db *sql.DB) iamDomain.UserRepository {
	return &userRepository{
		db:      db,
		queries: database.New(db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *iamDomain.User) (*iamDomain.User, error) {
	dbUser, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Email:        string(user.Email),
		PasswordHash: string(user.PasswordHash),
		FullName:     user.FullName,
	})
	if err != nil {
		return nil, err
	}

	return toDomainUser(dbUser)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email iamDomain.Email) (*iamDomain.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, string(email))
	if err != nil {
		return nil, err
	}

	return toDomainUser(user)
}

func toDomainUser(dbUser database.User) (*iamDomain.User, error) {
	email, err := iamDomain.NewEmail(dbUser.Email)
	if err != nil {
		return nil, err
	}

	return &iamDomain.User{
		ID:           dbUser.ID,
		Email:        email,
		PasswordHash: iamDomain.PasswordHash(dbUser.PasswordHash),
		FullName:     dbUser.FullName,
	}, nil
}
