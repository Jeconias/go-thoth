package models

// DirValidate TODO
type DirValidate struct {
	Dir     string   `thoth:"dir"`
	Pointer *string  `thoth:"dir"`
	Slice   []string `thoth:"dir"`
}
