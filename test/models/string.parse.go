package models

// TypeEqString TODO
type TypeEqString struct {
	String  string  `thoth:"eq=chico"`
	Pointer *string `thoth:"eq=bento"`
}
