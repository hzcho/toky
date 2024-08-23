package usecase

import(
	"context"
)
type Auth interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	VerifyToken(ctx context.Context, token string) (int, error)
}
