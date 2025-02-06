package controllers

import (
	"net/http"
	"nganterin-go/api/storages/services"
	"nganterin-go/api/storages/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	file, fileErr := ctx.FormFile("file")
	if fileErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	if file.Size > (10 * 1024 * 1024) {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrFileSize))
		return
	}

	fileContent, fileErr := file.Open()
	if fileErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrFileRead))
		return
	}
	defer fileContent.Close()

	buffer := make([]byte, file.Size)
	_, fileErr = fileContent.Read(buffer)
	if fileErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrFileRead))
		return
	}

	fileName := file.Filename
	fileExtension := fileName[strings.LastIndex(fileName, ".")+1:]
	mimeType := file.Header.Get("Content-Type")
	mimeParts := strings.Split(mimeType, "/")
	mimeMainType, mimeSubType := mimeParts[0], ""
	if len(mimeParts) > 1 {
		mimeSubType = mimeParts[1]
	}

	fileData := dto.FilesInputDTO{
		OriginalFileName: fileName,
		Size:             helpers.FormatFileSize(file.Size),
		Extension:        fileExtension,
		MimeType:         mimeMainType,
		MimeSubType:      mimeSubType,
		Meta:             "{}",
	}

	result, err := h.services.Create(ctx, buffer, fileData)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "file uploaded successfully",
		Data:    result,
	})
}
