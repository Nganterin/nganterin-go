package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"nganterin-go/api/storages/repositories"
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"
	"nganterin-go/pkg/mapper"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
	DB   *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		repo: compRepositories,
		DB:   db,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, file []byte, data dto.FilesInputDTO) (*dto.FilesOutputDTO, *exceptions.Exception) {
	modelData := mapper.MapFilesInputToModel(data)

	uniqueName := helpers.GenerateUniqueFileName() + "." + data.Extension

	publicURL, metadata, err := s.DriveUpload(ctx, file, uniqueName, data.MimeType)
	if err != nil {
		return nil, err
	}

	modelData.PublicURL = *publicURL
	modelData.Meta = *metadata

	repoData, err := s.repo.Create(ctx, s.DB, modelData)
	if err != nil {
		return nil, err
	}

	modelData.ID = repoData.ID
	result := mapper.MapFilesModelToOutput(*repoData)

	return &result, nil
}

func (s *CompServicesImpl) DriveUpload(ctx *gin.Context, file []byte, name, mimeType string) (*string, *string, *exceptions.Exception) {
	APPLICATION_CREDENTIALS := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	APPLICATION_FOLDER_ID := os.Getenv("APPLICATION_FOLDER_ID")

	driveService, err := drive.NewService(ctx, option.WithCredentialsJSON([]byte(APPLICATION_CREDENTIALS)))
	if err != nil {
		return nil, nil, exceptions.NewException(http.StatusBadGateway, exceptions.ErrFileUpload)
	}

	fileMetadata := &drive.File{
		Name:     name,
		Parents:  []string{APPLICATION_FOLDER_ID},
		MimeType: mimeType,
	}

	fileReader := bytes.NewReader(file)
	uploadedFile, err := driveService.Files.Create(fileMetadata).
		Media(fileReader).
		Fields("id, name, mimeType, size, createdTime").
		Do()
	if err != nil {
		return nil, nil, exceptions.NewException(http.StatusBadGateway, exceptions.ErrFileUpload)
	}

	_, err = driveService.Permissions.Create(uploadedFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do()
	if err != nil {
		return nil, nil, exceptions.NewException(http.StatusBadGateway, exceptions.ErrFilePermission)
	}

	publicLink := fmt.Sprintf("https://lh3.googleusercontent.com/d/%s", uploadedFile.Id)

	metadata := map[string]interface{}{
		"id":          uploadedFile.Id,
		"name":        uploadedFile.Name,
		"mimeType":    uploadedFile.MimeType,
		"size":        uploadedFile.Size,
		"createdTime": uploadedFile.CreatedTime,
		"publicLink":  publicLink,
	}

	stringifiedMetadata, err := json.Marshal(metadata)
	if err != nil {
		return nil, nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrJsonMarshal)
	}

	return &publicLink, helpers.StringPointer(stringifiedMetadata), nil
}
