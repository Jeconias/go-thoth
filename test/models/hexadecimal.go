package models

// HexadecimalValidate TODO
type HexadecimalValidate struct {
	Hexadecimal string   `thoth:"hexadecimal"`
	Pointer     *string  `thoth:"hexadecimal"`
	Slice       []string `thoth:"hexadecimal"`
}
