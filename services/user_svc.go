package services

import (
	"errors"
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

	_, err = s.repo.RegisterUserCredential(user_data)
	if err != nil {
		return err
	}

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

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["name"] = data.Name

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
