package usecase

import (
	"context"
	"io"
	"toky/internal/domain/model"
)

type FileUseCase interface{
	UploadFile(ctx context.Context, metadata *model.FileMetadata, file io.Reader) error
}