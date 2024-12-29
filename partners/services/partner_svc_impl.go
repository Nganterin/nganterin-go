package services

import (
	"fmt"
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"
	"nganterin-go/partners/repositories"
	"os"
	"time"

	emailServices "nganterin-go/emails/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Partner) *exceptions.Exception {
	hashed_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		return err
	}

	partner_data := database.Partners{
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

	partner_id, err := s.repo.Create(ctx, s.DB, partner_data)
	if err != nil {
		return err
	}

	token, err := s.repo.CreateVerificationToken(ctx, s.DB, *partner_id)
	if err != nil {
		return err
	}

	go func() {
		verificationEmail := dto.EmailVerification{
			Email:           data.Email,
			Subject:         "Nganterin - Verification Partner Email",
			VerificationURL: os.Getenv("DASHBOARD_BASE_URL") + "/auth/verify?token=" + *token,
		}

		err = emailServices.NewComponentServices().VerificationEmail(verificationEmail)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func (s *CompServicesImpl) Login(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception) {
	data, err := s.repo.FindByEmail(ctx, s.DB, email)
	if err != nil {
		return nil, err
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
	claims["name"] = data.Name
	claims["email"] = data.Email
	claims["company_name"] = data.CompanyName
	claims["owner"] = data.Owner
	claims["company_field"] = data.CompanyField
	claims["company_email"] = data.CompanyEmail
	claims["company_address"] = data.CompanyAddress
	claims["is_partner"] = true

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}

func (s *CompServicesImpl) VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception {
	return s.repo.VerifyEmail(ctx, s.DB, token)
}
