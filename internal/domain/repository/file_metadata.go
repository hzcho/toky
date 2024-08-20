package repositories

import (
    "context"
    "toky/internal/domain/model"
)

type FileMetadataRepository interface {
    Save(ctx context.Context, metadata *model.FileMetadata) error
    GetByID(ctx context.Context, id string) (*model.FileMetadata, error)
    DeleteByID(ctx context.Context, id string) error
    Update(ctx context.Context, metadata *model.FileMetadata) error
    List(ctx context.Context) ([]*model.FileMetadata, error)
}
