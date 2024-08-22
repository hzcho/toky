package repository

import (
	"context"
	"io"
)

type FileStorageRepository interface {
	SaveFile(ctx context.Context, fileName string, file io.Reader) error
	GetFile(ctx context.Context, fileName string) (io.ReadCloser, error)
	DeleteFile(ctx context.Context, fileName string) error
}
