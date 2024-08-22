package handler

import (
	"github.com/labstack/echo/v4"
	"toky/internal/domain/usecase"
)

type Handler struct {
	*FileGroup
}

func New(e *echo.Echo, usecase usecase.FileUseCase) *Handler {
	fileGroup := NewFileGroup(&usecase)

	handler := &Handler{
		FileGroup: fileGroup,
	}

	auth:=e.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}
	api := e.Group("/api/v1")
	{
		files := api.Group("/files")
		{
			files.POST("/", handler.Save)
		}
	}

	return handler
}
