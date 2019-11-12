package models

// NeField TODO
type NeField struct {
	ID     int64  `thoth:"-"`
	Status bool   `thoth:"required"`
	Name   string `thoth:"ne=chico"`
}

// NeFieldStrPointer TODO
type NeFieldStrPointer struct {
	ID     int64   `thoth:"-"`
	Status bool    `thoth:"required"`
	Name   *string `thoth:"ne=bento"`
}
