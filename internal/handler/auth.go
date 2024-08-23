package handler

import (
	"net/http"
	"toky/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

type AuthGroup struct{
	auth usecase.Auth
}

type AuthReq struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func (g *AuthGroup) SignIn(c echo.Context) error{
	input:=AuthReq{}

	if err:=c.Bind(&input); err!=nil{
		return c.JSON(http.StatusBadRequest, "incorrect request structure")
	}

	token, err:=g.auth.GenerateToken(c.Request().Context(), input.Email, input.Password)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "couldn't create the token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (g *AuthGroup) Register(c echo.Context) error{
	input:=AuthReq{}

	if err:=c.Bind(&input); err!=nil{
		return c.JSON(http.StatusBadRequest, "incorrect request structure")
	}

	userId, err:=g.auth.CreateUser(c.Request().Context(), input.Email, input.Password)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "couldn't create the user")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userId": userId,
	})
}