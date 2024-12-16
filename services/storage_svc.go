package services

import (
	"bytes"
	"context"
	"errors"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func (s *compServices) SaveFileToDrive(file []byte, name string, format string) (string, error) {
	ctx := context.Background()
		driveService, err := drive.NewService(ctx, option.WithCredentialsFile("secret.json"))
	if err != nil {
		return "", errors.New("failed to create drive service: " + err.Error())
	}

	fileMetadata := &drive.File{
		Name:    name,
		MimeType: format,
	}

	fileReader := bytes.NewReader(file)

	uploadedFile, err := driveService.Files.Create(fileMetadata).
		Media(fileReader).
		Do()
	if err != nil {
		return "", errors.New("failed to upload file to drive: " + err.Error())
	}

	_, err = driveService.Permissions.Create(uploadedFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do()
	if err != nil {
		return "", errors.New("failed to set file permissions: " + err.Error())
	}

	publicLink := "https://drive.google.com/file/d/" + uploadedFile.Id + "/view?usp=sharing"

	return publicLink, nil
}
