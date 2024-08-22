package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
	"toky/internal/domain/model"
	"toky/internal/domain/repository"
	"path/filepath"
)

type File struct {
	StorageRepo  repository.FileStorageRepository
	MetadataRepo repository.FileMetadataRepository
}

func (u *File) UploadFile(ctx context.Context, metadata *model.FileMetadata, file io.Reader) error {
    basePath := filepath.Join("uploads", metadata.Path)
    path := basePath

    for i := 0; ; i++ {
        if _, err := os.Stat(path); os.IsNotExist(err) {
            metadata.Path = path
            break
        }

        path = fmt.Sprintf("%s_%d", basePath, i)
    }

    if err := u.StorageRepo.SaveFile(ctx, metadata.Path, file); err != nil {
        return err
    }

    if err := u.MetadataRepo.Save(ctx, metadata); err != nil {
        return err
    }

    return nil
}
