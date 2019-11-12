package models

// LtValidate TODO
type LtValidate struct {
	Name     string   `thoth:"lt=12"`
	Password *string  `thoth:"lt=3"`
	Age      int      `thoth:"lt=22"`
	Contents []string `thoth:"lt=2"`
}
