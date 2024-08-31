package repository

import (
	"toky/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)
type Repositories struct{
	repository.User
	repository.FileMetadata
	repository.FileStorage
}

func NewRepositories(db*pgxpool.Pool, uploadDir string)*Repositories{
	return &Repositories{
		User: NewUser(db),
		FileMetadata: NewFileMetadata(db),
		FileStorage: NewFileStorage(uploadDir),
	}
}