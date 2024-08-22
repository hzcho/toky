package handler

import (
	"net/http"
	"strings"
	"toky/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeader = "Authorization"
	userEmail           = "email"
)

type Middleware struct {
	auth usecase.Auth
}

func (m *Middleware) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return c.JSON(http.StatusUnauthorized, "No authorization header")
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "Invalid authorization header")
		}

		if len(parts[1]) == 0 {
			return c.JSON(http.StatusUnauthorized, "token is empty")
		}

		email, err := m.auth.VerifyToken(parts[1])
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "invalid token")
		}

		c.Set(userEmail, email)

		return nil
	}
}
