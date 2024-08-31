package usecase

import (
	"context"
	"io"
	"toky/internal/domain/model"
)

type File interface {
	UploadFile(ctx context.Context, metadata *model.FileMetadata, file io.Reader) error
}
