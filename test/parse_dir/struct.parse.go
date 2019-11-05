package any

// StructA TODO
type StructA struct {
	A string
	B int
}

// Struct TODO
type Struct struct {
	A StructA     `thoth:"required"`
	B *StructA    `thoth:"required"`
	C []StructA   `thoth:"required"`
	D []*StructA  `thoth:"required"`
	E *[]StructA  `thoth:"required"`
	F *[]*StructA `thoth:"required"`
}
