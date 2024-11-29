package dto

type Partner struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	CompanyName    string `json:"company_name" binding:"required"`
	Owner          string `json:"owner" binding:"required"`
	CompanyField   string `json:"company_field" binding:"required"`
	CompanyEmail   string `json:"company_email" binding:"required"`
	CompanyAddress string `json:"company_address" binding:"required"`
	LegalityFile   string `json:"legality_file" binding:"required"`
	MOUFile        string `json:"mou_file" binding:"required"`
}
