package models

// LteValidate TODO
type LteValidate struct {
	Name     string   `thoth:"lte=12"`
	Password *string  `thoth:"lte=3"`
	Age      int      `thoth:"lte=22"`
	Contents []string `thoth:"lte=2"`
}
