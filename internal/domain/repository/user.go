package repository

import (
	"toky/internal/domain/model"
	"context"
)

type User interface {
	User(ctx context.Context, name, passHash string) (model.User, error)
	Save(ctx context.Context, name, passHash string) (uint64, error)
}
