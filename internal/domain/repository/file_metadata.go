package repository

import (
    "context"
    "toky/internal/domain/model"
)

type FileMetadataRepository interface {
    Save(ctx context.Context, metadata *model.FileMetadata) error
    GetByName(ctx context.Context, fileName string) (*model.FileMetadata, error)
    DeleteByID(ctx context.Context, id string) error
    List(ctx context.Context) ([]*model.FileMetadata, error)
}
