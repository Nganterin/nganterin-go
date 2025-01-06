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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Partner) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

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

	partner_id, err := s.repo.Create(ctx, tx, partner_data)
	if err != nil {
		return err
	}

	token, err := s.repo.CreateVerificationToken(ctx, tx, *partner_id)
	if err != nil {
		return err
	}

	go func() {
		verificationEmail := dto.EmailVerification{
			Email:           data.Email,
			Subject:         "Nganterin - Verification Partner Email",
			VerificationURL: os.Getenv("DASHBOARD_BASE_URL") + "/auth/verify?token=" + *token,
		}

		err = s.emailService.VerificationEmail(verificationEmail)
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
	claims["is_data_verified"] = data.DataVerifiedAt != nil
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

func (s *CompServicesImpl) ApprovalCheck(ctx *gin.Context, id string) (*string, *exceptions.Exception) {
	data, err := s.repo.FindByID(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	if data.DataVerifiedAt == nil {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrDataNotVerified)
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
	claims["is_data_verified"] = data.DataVerifiedAt != nil
	claims["is_partner"] = true

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}