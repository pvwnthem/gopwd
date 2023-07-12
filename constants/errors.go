package constants

const (
	ErrVaultExistence       = "failed to check vault existence: %w"
	ErrVaultDoesExist       = "vault already exists at %s"
	ErrVaultDoesNotExist    = "vault does not exist at %s"
	ErrPasswordExistence    = "failed to check password existence: %w"
	ErrPasswordDoesExist    = "password already exists"
	ErrPasswordDoesNotExist = "password does not exist"
	ErrDirectoryCreation    = "failed to create directory: %w"
	ErrGetGPGID             = "failed to get gpg-id: %w"
	ErrPasswordEncryption   = "failed to encrypt password: %w"
	ErrPasswordWrite        = "failed to write password to file: %w"
	ErrActionConfirm        = "failed to confirm action: %w"
)
