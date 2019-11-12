package models

// GteField TODO
type GteField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"gte=10"`
}

// GteFieldStrPointer TODO
type GteFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"gte=12"`
}

// GteFields TODO
type GteFields struct {
	ID     int  `thoth:"eq=1"`
	Status bool `thoth:"required"`
}
