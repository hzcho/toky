package usecase

import (
	"context"
)

type Auth interface {
	CreateUser(ctx context.Context, email, password string) (uint64, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	VerifyToken(ctx context.Context, tkn string) (string, error)
}
