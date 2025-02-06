package dto

type EmailRequest struct {
	Email   string
	Subject string
	Body    string
}

type EmailVerification struct {
	Email           string
	Subject         string
	VerificationURL string
}
