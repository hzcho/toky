package usecase

import (
	"log/slog"
	"toky/internal/domain/usecase"
	"toky/internal/repository"
)

type Usecases struct {
	usecase.Auth
	usecase.File
}

func NewUsecases(repositories *repository.Repositories, log *slog.Logger, ttl int64) *Usecases {
	return &Usecases{
		Auth: NewAuth(repositories.User, log, ttl),
		File: NewFile(repositories.FileStorage, repositories.FileMetadata, log),
	}
}
