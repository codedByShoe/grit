package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	User      User   `gorm:"foreignKey:UserID" json:"user"`
	SessionID string `json:"session_id"`
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) Create(session *Session) (*Session, error) {
	session.SessionID = uuid.New().String()
	result := r.db.Create(session)

	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (r *SessionRepository) GetUserFromSession(sessionID string, userID string) (*User, error) {
	var session Session

	err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email")
	}).Where("session_id = ? AND user_id = ?", sessionID, userID).First(&session).Error
	if err != nil {
		return nil, err
	}

	if session.User.ID == 0 {
		return nil, fmt.Errorf("no user associated with the session")
	}

	return &session.User, nil
}
