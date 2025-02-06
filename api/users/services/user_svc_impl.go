package services

import (
	"fmt"
	"net/http"
	"nganterin-go/api/users/repositories"
	"nganterin-go/models"
	"nganterin-go/api/users/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"
	"os"
	"time"

	emailServices "nganterin-go/emails/services"
	emailDTO "nganterin-go/emails/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo         repositories.CompRepositories
	emailService emailServices.CompServices
	DB           *gorm.DB
	validate     *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, compEmailService emailServices.CompServices, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:         compRepositories,
		emailService: compEmailService,
		DB:           db,
		validate:     validate,
	}
}

func (s *CompServicesImpl) CreateCredentials(ctx *gin.Context, data dto.User) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hashed_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrCredentialsHash)
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

	userID, err := s.repo.Create(ctx, tx, user_data)
	if err != nil {
		return err
	}

	token, err := s.repo.CreateVerificationToken(ctx, tx, *userID)
	if err != nil {
		return err
	}

	go func() {
		verificationEmail := emailDTO.EmailVerification{
			Email:           data.Email,
			Subject:         "Nganterin - Verification Email",
			VerificationURL: os.Getenv("WEBCLIENT_BASE_URL") + "/auth/verify?token=" + *token,
		}

		err = s.emailService.VerificationEmail(verificationEmail)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func (s *CompServicesImpl) CreateGoogleOAuth(ctx *gin.Context, data dto.UserGoogle) (*string, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hashedGoogleSUB, err := helpers.HashPassword(data.GoogleSUB)
	if err != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrCredentialsHash)
	}

	now := time.Now()

	user_data := models.Users{
		Name:            data.Name,
		Email:           data.Email,
		EmailVerifiedAt: &now,
		HashedGoogleSUB: hashedGoogleSUB,
		PhoneNumber:     data.PhoneNumber,
		Country:         data.Country,
		Province:        data.Province,
		City:            data.City,
		ZipCode:         data.ZipCode,
		CompleteAddress: data.CompleteAddress,
		Avatar:          data.Avatar,
	}

	userID, err := s.repo.Create(ctx, tx, user_data)
	if err != nil {
		return nil, err
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userID
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["email_verified_at"] = data.EmailVerifiedAt
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress
	claims["avatar"] = data.Avatar

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signError := token.SignedString(secretKey)
	if signError != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}

func (s *CompServicesImpl) LoginCredentials(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception) {
	data, err := s.repo.FindByEmail(ctx, s.DB, email)
	if err != nil {
		return nil, err
	}

	if data.HashedGoogleSUB != "" {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrRegisteredWithGoogle)
	}

	err = helpers.CheckPasswordHash(password, data.HashedPassword)
	if err != nil {
		return nil, err
	}

	if data.EmailVerifiedAt == nil {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrEmailNotVerified)
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["email_verified_at"] = data.EmailVerifiedAt
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress
	claims["avatar"] = data.Avatar

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signError := token.SignedString(secretKey)
	if signError != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}

func (s *CompServicesImpl) VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.VerifyEmail(ctx, tx, token)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) LoginGoogleOAuth(ctx *gin.Context, email string, googleSUB string) (*string, *exceptions.Exception) {
	data, err := s.repo.FindByEmail(ctx, s.DB, email)
	if err != nil {
		return nil, err
	}

	if data.HashedPassword != "" {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrRegisteredWithCredentials)
	}

	err = helpers.CheckPasswordHash(googleSUB, data.HashedGoogleSUB)
	if err != nil {
		return nil, err
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["email_verified_at"] = data.EmailVerifiedAt
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress
	claims["avatar"] = data.Avatar

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}
