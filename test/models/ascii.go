package models

// ASCIIValidate TODO
type ASCIIValidate struct {
	ASCII   string   `thoth:"ascii"`
	Pointer *string  `thoth:"ascii"`
	Slice   []string `thoth:"ascii"`
}
