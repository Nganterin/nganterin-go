package services

import (
	"fmt"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"nganterin-go/models"
)

func (s *compServices) RegisterPartner(data dto.Partner) error {
	hashed_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		return err
	}

	partner_data := models.Partners{
		Name:           data.Name,
		Email:          data.Email,
		HashedPassword: hashed_password,
		CompanyName:    data.CompanyName,
		Owner:          data.Owner,
		CompanyField:   data.CompanyField,
		CompanyEmail:   data.CompanyEmail,
		CompanyAddress: data.CompanyAddress,
		LegalityFile:   data.LegalityFile,
		MOUFile:        data.MOUFile,
	}

	partner_id, err := s.repo.RegisterPartner(partner_data)
	if err != nil {
		return err
	}

	go func() {
		err := s.GeneratePartnerVerificationEmail(*partner_id)
		if err != nil {
			fmt.Printf("Error generating verification email: %v\n", err)
		}
	}()

	return nil
}
