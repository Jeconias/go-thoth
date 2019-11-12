package models

// MinField TODO
type MinField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"min=10"`
}

// MinFieldStrPointer TODO
type MinFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"min=12"`
}

// MinFields TODO
type MinFields struct {
	ID       int       `thoth:"eq=1"`
	Status   bool      `thoth:"required"`
	Names    []*string `thoth:"min=2"`
	Subjects *[]string `thoth:"min=3"`
}
