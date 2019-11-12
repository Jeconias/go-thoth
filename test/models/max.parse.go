package models

// MaxField TODO
type MaxField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"max=10"`
}

// MaxFieldStrPointer TODO
type MaxFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"max=12"`
}

// MaxFields TODO
type MaxFields struct {
	ID     int  `thoth:"eq=1"`
	Status bool `thoth:"required"`
}
