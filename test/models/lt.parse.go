package models

// LtField TODO
type LtField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"lt=10"`
}

// LtFieldStrPointer TODO
type LtFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"lt=12"`
}

// LtFields TODO
type LtFields struct {
	ID     int  `thoth:"eq=1"`
	Status bool `thoth:"required"`
}
