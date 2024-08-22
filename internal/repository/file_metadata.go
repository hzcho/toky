package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"toky/internal/domain/model"
)

type FileMetadata struct {
	db pgxpool.Pool
}

func (r *FileMetadata) Save(ctx context.Context, metadata *model.FileMetadata) error {
	query := "insert into file_metadata (file_name, path, created_at) values ($1, $2, $3, $4)"

	_, err := r.db.Exec(ctx, query, metadata.FileName, metadata.Path, metadata.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileMetadata) GetByName(ctx context.Context, fileName string) (*model.FileMetadata, error) {
	var metadata model.FileMetadata

	query := "select (id, file_name, path, size, created_at) from file_metadata where file_name=$1"

	err := r.db.QueryRow(ctx, query, fileName).Scan(
		&metadata.ID,
		&metadata.FileName,
		&metadata.Path,
		&metadata.Size,
		&metadata.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (r *FileMetadata) DeleteByID(ctx context.Context, id string) error {
	query := "delete from file_metadata where id=$1"

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (r *FileMetadata) List(ctx context.Context) ([]*model.FileMetadata, error) {
	query := "select * from file_metadata"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metadataList []*model.FileMetadata

	for rows.Next() {
		metadata := &model.FileMetadata{}

		err := rows.Scan(
			&metadata.ID,
			&metadata.FileName,
			&metadata.Path,
			&metadata.Size,
			&metadata.CreatedAt,
		)
		if err!=nil {
			return nil, err
		}

		metadataList=append(metadataList, metadata)
	}

	return metadataList, nil
}
