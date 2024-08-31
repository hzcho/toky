package repository

import (
	"context"
	"toky/internal/domain/model"
)

type User interface {
	User(ctx context.Context, email string) (model.User, error)
	Save(ctx context.Context, email, passHash string) (uint64, error)
}
