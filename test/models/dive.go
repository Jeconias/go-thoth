package models

// DiveValidate TODO
type DiveValidate struct {
	SliceString       []string          `thoth:"dive"`
	SliceInt          []int             `thoth:"dive"`
	MapStringToString map[string]string `thoth:"dive"`
}
