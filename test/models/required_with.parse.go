package models

// RequiredWithField TODO
type RequiredWithField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"required_with=Status"`
}

// RequiredWithFieldStrPointer TODO
type RequiredWithFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_with=Status"`
}

// RequiredWithFields TODO
type RequiredWithFields struct {
	ID     int     `thoth:"eq=1"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_with=ID Status"`
}
