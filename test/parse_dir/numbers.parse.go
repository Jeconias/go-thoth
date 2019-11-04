package any

// Number TODO
type Number struct {
	//  uint is an unsigned integer type
	A uint    `thoth:"required"`
	B uint8   `thoth:"required"`
	C uint16  `thoth:"required"`
	D uint32  `thoth:"required"`
	E uint64  `thoth:"required"`
	F uintptr `thoth:"required"`

	// int is a signed integer type
	G int   `thoth:"required"`
	H int8  `thoth:"required"`
	I int16 `thoth:"required"`
	J int32 `thoth:"required"`
	K int64 `thoth:"required"`

	// float is the set of all IEEE-754
	L float32 `thoth:"required"`
	M float64 `thoth:"required"`

	// complex is the set of all complex numbers with float32 real and imaginary parts.
	N complex64  `thoth:"required"`
	O complex128 `thoth:"required"`
}
