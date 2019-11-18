package models

// RequiredWithAllField TODO
type RequiredWithAllField struct {
	Status bool   `thoth:"required"`
	Name   string `thoth:"required_with_all=Status"`
}

// RequiredWithAllFieldStrPointer TODO
type RequiredWithAllFieldStrPointer struct {
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_with_all=Status"`
}

// RequiredWithAllFields TODO
type RequiredWithAllFields struct {
	ID     int     `thoth:"eq=1"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_with_all=ID Status"`
}
