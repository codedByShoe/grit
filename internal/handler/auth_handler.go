package handler

import (
	"github.com/codedbyshoe/grit/internal/model"
	"github.com/codedbyshoe/grit/internal/service"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*chi.Mux
	userRepo          model.UserRepository
	sessionRepo       model.SessionRepository
	passwordService   service.PasswordHash
	db                *gorm.DB
	sessionCookieName string
}

type AuthHandlerParams struct {
	userRepo          model.UserRepository
	sessionRepo       model.SessionRepository
	passwordService   service.PasswordHash
	db                *gorm.DB
	sessionCookieName string
}

func NewAuthHandler(params AuthHandlerParams) *AuthHandler {
	h := &AuthHandler{
		Mux:               chi.NewRouter(),
		userRepo:          params.userRepo,
		sessionRepo:       params.sessionRepo,
		passwordService:   params.passwordService,
		db:                params.db,
		sessionCookieName: params.sessionCookieName,
	}

	return h
}
