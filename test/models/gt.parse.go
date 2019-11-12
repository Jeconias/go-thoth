package models

// GtField TODO
type GtField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"gt=10"`
}

// GtFieldStrPointer TODO
type GtFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"gt=12"`
}

// GtFields TODO
type GtFields struct {
	ID     int  `thoth:"eq=1"`
	Status bool `thoth:"required"`
}
