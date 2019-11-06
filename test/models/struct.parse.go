package models

// StructA TODO
type StructA struct {
	A string
	B int
}

// Struct TODO
type Struct struct {
	Struct                    StructA     `thoth:"required"`
	StructPointer             *StructA    `thoth:"required"`
	SliceStruct               []StructA   `thoth:"required"`
	SliceStructPointer        []*StructA  `thoth:"required"`
	SlicePointerStruct        *[]StructA  `thoth:"required"`
	SlicePointerStructPointer *[]*StructA `thoth:"required"`
}
