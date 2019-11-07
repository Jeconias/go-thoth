package models

// TypeString TODO
type TypeString struct {
	String  string  `thoth:"required"`
	Pointer *string `thoth:"required"`
}

// TypeEqString TODO
type TypeEqString struct {
	String  string  `thoth:"eq=chico"`
	Pointer *string `thoth:"eq=bento"`
}
