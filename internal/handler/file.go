package handler

import (
	"net/http"
	"toky/internal/domain/model"
	"toky/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

type FileSaveReq struct {
	FileName string `form:"filename" json:"filename"`
	File     []byte `form:"file" json:"file"`
}

type FileGroup struct {
	fileUseCase usecase.FileUseCase
}

func NewFileGroup(usecase usecase.FileUseCase) *FileGroup{
	return &FileGroup{
		fileUseCase: usecase,
	}
}

func (g *FileGroup) Save(c echo.Context) error {
	req := new(FileSaveReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect file")
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal error")
	}
	defer src.Close()

	metadata := &model.FileMetadata{
		FileName: file.Filename,
		Size:     file.Size,
	}

	if err := g.fileUseCase.UploadFile(c.Request().Context(), metadata, src); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to upload file")
	}

	return c.JSON(http.StatusOK, "file uploaded successfully")
}