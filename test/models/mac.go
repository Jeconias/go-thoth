package models

// MACValidate TODO
type MACValidate struct {
	MAC     string   `thoth:"mac"`
	Pointer *string  `thoth:"mac"`
	Slice   []string `thoth:"mac"`
}
