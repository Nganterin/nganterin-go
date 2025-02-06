package services

import (
	"nganterin-go/emails/dto"
	"nganterin-go/pkg/exceptions"
)

type CompServices interface {
	SendEmail(data dto.EmailRequest) *exceptions.Exception
	VerificationEmail(data dto.EmailVerification) *exceptions.Exception
}
