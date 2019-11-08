package models

// RequiredWith TODO
type RequiredWith struct {
	ID   int64
	Name string `thoth:"required_with=Age"`
	Age  bool   `thoth:"required"`
}
