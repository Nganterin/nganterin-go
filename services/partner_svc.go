package services

import (
	"errors"
	"fmt"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"nganterin-go/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (s *compServices) LoginPartner(email string, password string) (*string, error) {
	data, err := s.repo.GetPartnerDetailsByEmail(email)
	if err != nil {
		return nil, errors.New("404")
	}

	if !helpers.CheckPasswordHash(password, data.HashedPassword) {
		return nil, errors.New("401")
	}

	if data.EmailVerifiedAt == nil {
		return nil, errors.New("403")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["name"] = data.Name
	claims["email"] = data.Email
	claims["company_name"]= data.CompanyName
	claims["owner"] = data.Owner
	claims["company_field"] = data.CompanyField
	claims["company_email"] = data.CompanyEmail
	claims["company_address"] = data.CompanyAddress

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}