package service

import (
	"github.com/codedbyshoe/grit/internal/model"
	"github.com/codedbyshoe/grit/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.FindByEmail(email)
}
