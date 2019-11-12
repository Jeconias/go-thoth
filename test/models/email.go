package models

// EmailValidation TODO
type EmailValidation struct {
	Email   string  `thoth:"email"`
	Pointer *string `thoth:"email"`
}
