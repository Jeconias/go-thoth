package models

// Number TODO
type Number struct {
	//  uint is an unsigned integer type
	Uint           uint     `thoth:"required"`
	UintPointer    *uint    `thoth:"required"`
	Uint8          uint8    `thoth:"required"`
	Uint8Pointer   *uint8   `thoth:"required"`
	Uint16         uint16   `thoth:"required"`
	Uint16Pointer  *uint16  `thoth:"required"`
	Uint32         uint32   `thoth:"required"`
	Uint32Pointer  *uint32  `thoth:"required"`
	Uint64         uint64   `thoth:"required"`
	Uint64Pointer  *uint64  `thoth:"required"`
	Uintptr        uintptr  `thoth:"required"`
	UintptrPointer *uintptr `thoth:"required"`

	// int is a signed integer type
	Int          int    `thoth:"required"`
	IntPointer   *int   `thoth:"required"`
	Int8         int8   `thoth:"required"`
	Int8Pointer  *int8  `thoth:"required"`
	Int16        int16  `thoth:"required"`
	Int16Pointer *int16 `thoth:"required"`
	Int32        int32  `thoth:"required"`
	Int32Pointer *int32 `thoth:"required"`
	Int64        int64  `thoth:"required"`
	Int64Pointer *int64 `thoth:"required"`

	// float is the set of all IEEE-754
	Float32        float32  `thoth:"required"`
	Float32Pointer *float32 `thoth:"required"`
	Float64        float64  `thoth:"required"`
	Float64Pointer *float64 `thoth:"required"`

	// complex is the set of all complex numbers with float32 real and imaginary parts.
	Complex64         complex64   `thoth:"required"`
	Complex64Pointer  *complex64  `thoth:"required"`
	Complex128        complex128  `thoth:"required"`
	Complex128Pointer *complex128 `thoth:"required"`
}
