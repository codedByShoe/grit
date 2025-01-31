package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/codedbyshoe/grit/internal/model"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	sessionRepo       model.SessionRepository
	sessionCookieName string
}

func NewAuthMiddleware(r model.SessionRepository, c string) *AuthMiddleware {
	return &AuthMiddleware{
		sessionRepo:       r,
		sessionCookieName: c,
	}
}

type UserContextKey string

var UserKey UserContextKey = "user"

// RequireGuest ensures that the user is not authenticated (guest)
func (m *AuthMiddleware) RequireGuest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := m.hasValidCookie(c)
		if ok {
			return c.Redirect("/", fiber.StatusSeeOther)
		}
		return c.Next()
	}
}

// RequireAuth ensures that the user is authenticated
func (m *AuthMiddleware) RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := m.hasValidCookie(c)
		if !ok {
			return c.Redirect("/auth", fiber.StatusSeeOther)
		}
		return c.Next()
	}
}

// AddUserToContext adds the authenticated user to the request context
func (m *AuthMiddleware) AddUserToContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := m.hasValidCookie(c)
		if !ok {
			return c.Next()
		}

		// Set the user in Fiber's context
		c.Locals(string(UserKey), user)
		return c.Next()
	}
}

// hasValidCookie checks if the session cookie is valid and returns the associated user
func (m *AuthMiddleware) hasValidCookie(c *fiber.Ctx) (*model.User, bool) {
	sessionCookie := c.Cookies(m.sessionCookieName)
	if sessionCookie == "" {
		return nil, false
	}

	decodedValue, err := base64.StdEncoding.DecodeString(sessionCookie)
	if err != nil {
		return nil, false
	}

	splitValue := strings.Split(string(decodedValue), ":")
	if len(splitValue) != 2 {
		return nil, false
	}

	sessionID := splitValue[0]
	userID := splitValue[1]

	user, err := m.sessionRepo.GetUserFromSession(sessionID, userID)
	if err != nil {
		return nil, false
	}
	return user, true
}

// GetUser retrieves the authenticated user from the context
func GetUser(c *fiber.Ctx) *model.User {
	user := c.Locals(string(UserKey))
	if user == nil {
		return nil
	}
	return user.(*model.User)
}
