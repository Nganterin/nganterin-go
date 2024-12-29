package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"
)

type CompServices interface {
	SendEmail(data dto.EmailRequest) *exceptions.Exception
	VerificationEmail(data dto.EmailVerification) *exceptions.Exception
}