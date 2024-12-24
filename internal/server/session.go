package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store       *sessions.CookieStore
	SessionName = "grit-session"
)

func NewStore(secret string) *sessions.CookieStore {

	store = sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return store
}

type Flash struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func SetFlash(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, flashType string, message string) {
	session, _ := store.Get(r, "flash")
	flash := Flash{
		Type:    flashType,
		Message: message,
	}

	jsonData, _ := json.Marshal(flash)
	session.AddFlash(jsonData)
	session.Save(r, w)
}
