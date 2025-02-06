package services

import (
	"nganterin-go/api/storages/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, file []byte, data dto.FilesInputDTO) (*dto.FilesOutputDTO, *exceptions.Exception)
	DriveUpload(ctx *gin.Context, file []byte, name, mimeType string) (*string, *string, *exceptions.Exception)
}
