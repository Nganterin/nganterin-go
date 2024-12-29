package mapper

import (
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

	"github.com/go-viper/mapstructure/v2"
)

func MapFilesInputToModel(input dto.FilesInputDTO) database.Files {
	var data database.Files
	mapstructure.Decode(input, &data)
	return data
}

func MapFilesModelToOutput(model database.Files) dto.FilesOutputDTO {
	var filesOutput dto.FilesOutputDTO
	mapstructure.Decode(model, &filesOutput)
	return filesOutput
}