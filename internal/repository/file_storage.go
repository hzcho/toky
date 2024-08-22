package repository

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type FileStorage struct{
	UploadDir string
}

func (r*FileStorage) SaveFile(ctx context.Context, fileName string, file io.Reader) error{
	fullPath:=filepath.Join(r.UploadDir, fileName)

	outFile, err:=os.Create(fullPath)
	if err!=nil {
		return err
	}
	defer outFile.Close()

	if _, err=io.Copy(outFile, file); err!=nil{
		return err
	}

	return nil
}

func (r*FileStorage) GetFile(ctx context.Context, fileName string) (io.ReadCloser, error){
	fullPath:=filepath.Join(r.UploadDir, fileName)

	file, err:=os.Open(fullPath)
	if err!=nil {
		return nil, err
	}

	return file, nil
}

func (r*FileStorage) DeleteFile(ctx context.Context, fileName string) error{
	fullPath:=filepath.Join(r.UploadDir, fileName)

	if err:=os.Remove(fullPath); err!=nil{
		return err
	}

	return nil
}