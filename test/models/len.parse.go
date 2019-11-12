package models

// LenField TODO
type LenField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"len=10"`
}

// LenFieldStrPointer TODO
type LenFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"len=12"`
}

// LenFields TODO
type LenFields struct {
	ID       int       `thoth:"eq=1"`
	Status   bool      `thoth:"required"`
	Names    []*string `thoth:"len=2"`
	Subjects *[]string `thoth:"len=3"`
}
