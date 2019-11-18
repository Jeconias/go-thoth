package models

// MaxValidate TODO
type MaxValidate struct {
	Name       string                 `thoth:"max=10"`
	Password   *string                `thoth:"max=8"`
	Age        int                    `thoth:"max=50"`
	Contents   []string               `thoth:"max=2"`
	Attributes map[string]interface{} `thoth:"max=3"`
}
