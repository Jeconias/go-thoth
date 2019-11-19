package models

// Base64Validate TODO
type Base64Validate struct {
	Base64  string   `thoth:"base64"`
	Pointer *string  `thoth:"base64"`
	Slice   []string `thoth:"base64"`
}
