package models

// EmailValidate TODO
type EmailValidate struct {
	Email   string  `thoth:"email"`
	Pointer *string `thoth:"email"`
}
