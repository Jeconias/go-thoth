package models

// Base64urlValidate TODO
type Base64urlValidate struct {
	Base64url string   `thoth:"base64url"`
	Pointer   *string  `thoth:"base64url"`
	Slice     []string `thoth:"base64url"`
}
