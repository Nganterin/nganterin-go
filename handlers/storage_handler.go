package handlers

import (
	"net/http"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status: http.StatusBadRequest,
			Error: "file can't be null",
		})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status: http.StatusBadRequest,
			Error: "error to open the uploaded file",
		})
		return
	}
	defer fileContent.Close()

	buffer := make([]byte, file.Size)
	_, err = fileContent.Read(buffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status: http.StatusBadRequest,
			Error: "failed to read the uploaded file content",
		})
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

	result, err := h.service.FileUpload(buffer, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status: http.StatusOK,
		Message: "file uploaded successfully",
		Data:    result,
	})
}
