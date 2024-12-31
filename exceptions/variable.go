package exceptions

const (
	ErrInvalidCredentials        = "Invalid credentials"
	ErrUnauthorized              = "Unauthorized access"
	ErrBadRequest                = "Invalid request body or parameters"
	ErrNotFound                  = "Record not found"
	ErrInternalServer            = "Something went wrong"
	ErrEmailNotVerified          = "Email not verified"
	ErrEmailSendFailed           = "Failed to send email"
	ErrEmailAlreadyRegistered    = "Email already registered"
	ErrDatabaseCommunication     = "Failed to communicate with database"
	ErrTokenGenerate             = "Failed to generate token"
	ErrCredentialsHash           = "Failed to secure credentials"
	ErrFileUpload                = "Failed to upload file"
	ErrFilePermission            = "Failed to set file permission"
	ErrFileSize                  = "File size exceeds the limit"
	ErrJsonMarshal               = "Failed to marshal JSON"
	ErrFileRead                  = "Failed to read file"
	ErrFileURL                   = "Invalid file URL"
	ErrRegisteredWithGoogle      = "User already registered with Google"
	ErrRegisteredWithCredentials = "User already registered with credentials"
)
