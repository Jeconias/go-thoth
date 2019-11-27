package models

// OmitemptyValidate TODO
type OmitemptyValidate struct {
	OmitEmptyString      string   `thoth:"omitempty,min=5,max=5"`
	OmitEmptyInt         int      `thoth:"omitempty,min=5,max=5"`
	OmitEmptySliceString []string `thoth:"omitempty,min=1"`
	OmitEmptySliceInt    []int    `thoth:"omitempty"`
}
