package models

// RequiredWithoutAllField TODO
type RequiredWithoutAllField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"required_without_all=Status"`
}

// RequiredWithoutAllFieldStrPointer TODO
type RequiredWithoutAllFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without_all=Status"`
}

// RequiredWithoutAllFields TODO
type RequiredWithoutAllFields struct {
	ID     int     `thoth:"eq=1"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without_all=ID Status"`
}
