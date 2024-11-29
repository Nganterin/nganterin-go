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

func (s *compServices) RegisterUserCredential(data dto.User) error {
	hashed_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		return err
	}

	user_data := models.Users{
		Name:            data.Name,
		Email:           data.Email,
		HashedPassword:  hashed_password,
		PhoneNumber:     data.PhoneNumber,
		Country:         data.Country,
		Province:        data.Province,
		City:            data.City,
		ZipCode:         data.ZipCode,
		CompleteAddress: data.CompleteAddress,
	}

	user_id, err := s.repo.RegisterUserCredential(user_data)
	if err != nil {
		return err
	}

	go func() {
		err = s.GenerateVerificationEmail(*user_id)
		if err != nil {
			fmt.Printf("Error generating verification email: %v\n", err)
		}
	}()

	return nil
}

func (s *compServices) LoginUserCredentials(email string, password string) (*string, error) {
	data, err := s.repo.GetUserDetailsByEmail(email)
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
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (s *compServices) VerifyUserEmail(token string) error {
	return s.repo.VerifyUserEmail(token)
}
