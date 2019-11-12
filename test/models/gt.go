package models

// GtValidate TODO
type GtValidate struct {
	Name     string   `thoth:"gt=5"`
	Password *string  `thoth:"gt=3"`
	Age      int      `thoth:"gt=22"`
	Contents []string `thoth:"gt=2"`
}
