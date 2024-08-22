package model

import "time"

type FileMetadata struct{
	ID string
	FileName string
	Path string
	Size int64
	CreatedAt time.Time
}