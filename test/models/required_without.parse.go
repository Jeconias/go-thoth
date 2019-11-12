package models

// RequiredWithoutField TODO
type RequiredWithoutField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"required_without=Status"`
}

// RequiredWithoutFieldStrPointer TODO
type RequiredWithoutFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without=Status"`
}

// RequiredWithoutFields TODO
type RequiredWithoutFields struct {
	ID     int     `thoth:"eq=1"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"required_without=ID Status"`
}
