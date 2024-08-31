package usecase

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"toky/internal/domain/model"
	"toky/internal/domain/repository"
)

type File struct {
	storageRepo  repository.FileStorage
	metadataRepo repository.FileMetadata
	log          *slog.Logger
}

func NewFile(storageRepo repository.FileStorage, metadataRepo repository.FileMetadata, log *slog.Logger) *File {
	return &File{
		storageRepo:  storageRepo,
		metadataRepo: metadataRepo,
		log:          log,
	}
}

func (u *File) UploadFile(ctx context.Context, metadata *model.FileMetadata, file io.Reader) error {
	const op = "internal/usecase/file/UploadFile"
	log := u.log.With(
		slog.String("operation", op),
	)

	name := metadata.FileName

	for i := 0; ; i++ {
		isExists, err := u.storageRepo.IsExists(name)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		if !isExists {
			break
		}

		name = fmt.Sprintf("%s_%d", metadata.FileName, i)
	}

	metadata.Path = filepath.Join(u.storageRepo.AbsPath(), name)

	if err := u.storageRepo.SaveFile(ctx, name, file); err != nil {
		log.Error(err.Error())
		return err
	}

	if err := u.metadataRepo.Save(ctx, metadata); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
