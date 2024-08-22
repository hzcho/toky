package usecase

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"toky/internal/domain/model"
	"toky/internal/domain/repository"
)

type File struct {
	storageRepo  repository.FileStorage
	metadataRepo repository.FileMetadata
    log *slog.Logger
}

func NewFile(storageRepo repository.FileStorage, metadataRepo repository.FileMetadata, log *slog.Logger)*File{
    return &File{
        storageRepo: storageRepo,
        metadataRepo: metadataRepo,
        log: log,
    }
}

func (u *File) UploadFile(ctx context.Context, metadata *model.FileMetadata, file io.Reader) error {
    const op="internal/usecase/file/UploadFile"
    log:=u.log.With(
        slog.String("operation", op),
    )

    basePath := filepath.Join("uploads", metadata.Path)
    path := basePath

    for i := 0; ; i++ {
        if _, err := os.Stat(path); os.IsNotExist(err) {
            log.Error(err.Error())

            metadata.Path = path
            break
        }

        path = fmt.Sprintf("%s_%d", basePath, i)
    }

    if err := u.storageRepo.SaveFile(ctx, metadata.Path, file); err != nil {
        log.Error(err.Error())
        return err
    }

    if err := u.metadataRepo.Save(ctx, metadata); err != nil {
        log.Error(err.Error())
        return err
    }

    return nil
}
