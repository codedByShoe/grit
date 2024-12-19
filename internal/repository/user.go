package repository

import (
	"github.com/codedbyshoe/grit/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{BaseRepository: NewBaseRepository[model.User](db)}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return r.FindOne(&model.User{Email: email})
}
