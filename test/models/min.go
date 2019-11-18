package models

// MinValidate TODO
type MinValidate struct {
	Name       string                 `thoth:"min=3"`
	Password   *string                `thoth:"min=3"`
	Age        int                    `thoth:"min=22"`
	Contents   []string               `thoth:"min=1"`
	Attributes map[string]interface{} `thoth:"min=2"`
}
