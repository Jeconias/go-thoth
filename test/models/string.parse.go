package models

// TypeString TODO
type TypeString struct {
	String  string  `thoth:"required"`
	Pointer *string `thoth:"required"`
}
