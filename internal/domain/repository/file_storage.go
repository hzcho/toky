package repositories

import (
	"context"
	"io"
)

type FileStorageRepository interface {
	SaveFile(ctx context.Context, path string, file io.Reader) error
	GetFile(ctx context.Context, path string) (io.ReadCloser, error)
	DeleteFile(ctx context.Context, path string) error
}
