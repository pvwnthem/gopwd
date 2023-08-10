package constants

const (
	ErrVaultExistence        = "failed to check vault existence: %w"
	ErrVaultDoesExist        = "vault already exists at %s"
	ErrVaultDoesNotExist     = "vault does not exist at %s"
	ErrPasswordExistence     = "failed to check password existence: %w"
	ErrPasswordDoesExist     = "password already exists"
	ErrPasswordDoesNotExist  = "password does not exist"
	ErrDirectoryCreation     = "failed to create directory: %w"
	ErrGetGPGID              = "failed to get gpg-id: %w"
	ErrPasswordEncryption    = "failed to encrypt password: %w"
	ErrPasswordDecryption    = "failed to decrypt password: %w"
	ErrPasswordWrite         = "failed to write password to file: %w"
	ErrActionConfirm         = "failed to confirm action: %w"
	ErrJSONUnmarshal         = "failed to unmarshal JSON: %w"
	ErrJSONMarshal           = "failed to marshal JSON: %w"
	ErrConfigExistence       = "failed to check config existence: %w"
	ErrConfigDoesNotExist    = "config does not exist at %s"
	ErrConfigDoesExist       = "config already exists"
	ErrConfigRead            = "failed to read config: %w"
	ErrConfigWrite           = "failed to write config: %w"
	ErrUnexpectedApiResponse = "unexpected API response: %v"
)
