package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type Partner struct {
	ID             string `json:"id"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	CompanyName    string `json:"company_name" validate:"required"`
	Owner          string `json:"owner" validate:"required"`
	CompanyField   string `json:"company_field" validate:"required"`
	CompanyEmail   string `json:"company_email" validate:"required,email"`
	CompanyAddress string `json:"company_address" validate:"required"`
	LegalityFile   string `json:"legality_file" validate:"required,url"`
	MOUFile        string `json:"mou_file" validate:"required,url"`
}
