package mapper

import (
	"nganterin-go/models"
	"nganterin-go/models/dto"

	"github.com/go-viper/mapstructure/v2"
)

func MapFilesInputToModel(input dto.FilesInputDTO) models.Files {
	var data models.Files
	mapstructure.Decode(input, &data)
	return data
}

func MapFilesModelToOutput(model models.Files) dto.FilesOutputDTO {
	var filesOutput dto.FilesOutputDTO
	mapstructure.Decode(model, &filesOutput)
	return filesOutput
}
