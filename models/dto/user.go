package dto

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password,omitempty" binding:"required"`
	EmailVerifiedAt string `json:"email_verified_at"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Country         string `json:"country" binding:"required"`
	Province        string `json:"province" binding:"required"`
	City            string `json:"city" binding:"required"`
	ZipCode         string `json:"zip_code" binding:"required"`
	CompleteAddress string `json:"complete_address" binding:"required"`
}

type UserGoogle struct {
	ID              string `json:"id"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	GoogleSUB       string `json:"google_sub" binding:"required"`
	EmailVerifiedAt string `json:"email_verified_at"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Country         string `json:"country" binding:"required"`
	Province        string `json:"province" binding:"required"`
	City            string `json:"city" binding:"required"`
	ZipCode         string `json:"zip_code" binding:"required"`
	CompleteAddress string `json:"complete_address" binding:"required"`
	Avatar          string `json:"avatar"`
}
