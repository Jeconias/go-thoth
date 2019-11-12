package models

// GteValidate TODO
type GteValidate struct {
	Name     string   `thoth:"gte=5"`
	Password *string  `thoth:"gte=3"`
	Age      int      `thoth:"gte=22"`
	Contents []string `thoth:"gte=2"`
}
