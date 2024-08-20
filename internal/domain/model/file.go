package model

import "time"

type FileMetadata struct{
	ID string
	FileName string
	ContentType string
	Path string
	CreatedAt time.Time
}