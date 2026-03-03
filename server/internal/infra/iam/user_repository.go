package postgres

import (
	"context"
	"database/sql"

	"github.com/shojib116/auditflow-api/internal/database"
	"github.com/shojib116/auditflow-api/internal/domain/iam"
)

type userRepository struct {
	db      *sql.DB
	queries *database.Queries
}

func NewUserRepository(db *sql.DB) iam.UserRepository {
	return &userRepository{
		db:      db,
		queries: database.New(db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *iam.User) (*iam.User, error) {
	dbUser, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Email:        string(user.Email),
		PasswordHash: user.PasswordHash,
		FullName:     user.FullName,
	})
	if err != nil {
		return nil, err
	}

	return toDomainUser(dbUser)
}

func toDomainUser(dbUser database.User) (*iam.User, error) {
	email, err := iam.NewEmail(dbUser.Email)
	if err != nil {
		return nil, err
	}

	return &iam.User{
		ID:           dbUser.ID,
		Email:        email,
		PasswordHash: dbUser.PasswordHash,
		FullName:     dbUser.FullName,
	}, nil
}
