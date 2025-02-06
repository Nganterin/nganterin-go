package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type FilesOutputDTO struct {
	ID               string `json:"id"` 
	PublicURL        string `json:"public_url"` 
	OriginalFileName string `json:"original_file_name"` 
	Size             string `json:"size"` 
	Extension        string `json:"extension"` 
	MimeType         string `json:"mime_type"` 
	MimeSubType      string `json:"mime_sub_type"`
	Meta             string `json:"meta"` 
	CreatedAt        string `json:"created_at"`
}