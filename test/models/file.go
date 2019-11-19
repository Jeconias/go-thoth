package models

// FileValidate TODO
type FileValidate struct {
	File    string   `thoth:"file"`
	Pointer *string  `thoth:"file"`
	Slice   []string `thoth:"file"`
}
