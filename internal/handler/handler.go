package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	fileGroup  *FileGroup
	authGroup  *AuthGroup
	middleware *Middleware
}

func NewHandler(e *echo.Echo, fileGroup *FileGroup, authGroup *AuthGroup, middleware *Middleware) *Handler {
	handler := &Handler{
		fileGroup:  fileGroup,
		authGroup:  authGroup,
		middleware: middleware,
	}

	auth := e.Group("/auth")
	{
		auth.POST("/login", handler.authGroup.SignIn)
		auth.POST("/register", handler.authGroup.Register)
	}
	api := e.Group("/api/v1", handler.middleware.userIdentity)
	{
		files := api.Group("/files")
		{
			files.POST("/", handler.fileGroup.Save)
		}
	}

	return handler
}
