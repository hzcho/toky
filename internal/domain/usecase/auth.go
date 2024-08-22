package usecase

import(
	"toky/internal/domain/model"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	VerifyToken(token string) (int, error)
}