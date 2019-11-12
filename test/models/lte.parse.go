package models

// LteField TODO
type LteField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"lte=10"`
}

// LteFieldStrPointer TODO
type LteFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"lte=12"`
}

// LteFields TODO
type LteFields struct {
	ID     int  `thoth:"eq=1"`
	Status bool `thoth:"required"`
}
