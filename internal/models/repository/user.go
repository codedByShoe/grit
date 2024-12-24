package repository

import (
	"context"

	"github.com/codedbyshoe/grit/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	fields := []string{"id", "email", "name", "password", "created_at"}

	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](db, "users", fields),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE email = $1"
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
