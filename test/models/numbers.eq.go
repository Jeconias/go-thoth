package models

// TypeEqNumber TODO
type TypeEqNumber struct {
	//  uint is an unsigned integer type
	Uint           uint     `thoth:"eq=1"`
	UintPointer    *uint    `thoth:"eq=2"`
	Uint8          uint8    `thoth:"eq=3"`
	Uint8Pointer   *uint8   `thoth:"eq=4"`
	Uint16         uint16   `thoth:"eq=5"`
	Uint16Pointer  *uint16  `thoth:"eq=6"`
	Uint32         uint32   `thoth:"eq=7"`
	Uint32Pointer  *uint32  `thoth:"eq=8"`
	Uint64         uint64   `thoth:"eq=9"`
	Uint64Pointer  *uint64  `thoth:"eq=10"`
	Uintptr        uintptr  `thoth:"eq=11"`
	UintptrPointer *uintptr `thoth:"eq=12"`

	// int is a signed integer type
	Int          int    `thoth:"eq=13"`
	IntPointer   *int   `thoth:"eq=14"`
	Int8         int8   `thoth:"eq=15"`
	Int8Pointer  *int8  `thoth:"eq=16"`
	Int16        int16  `thoth:"eq=17"`
	Int16Pointer *int16 `thoth:"eq=18"`
	Int32        int32  `thoth:"eq=19"`
	Int32Pointer *int32 `thoth:"eq=20"`
	Int64        int64  `thoth:"eq=21"`
	Int64Pointer *int64 `thoth:"eq=22"`

	// float is the set of all IEEE-754
	Float32        float32  `thoth:"eq=23"`
	Float32Pointer *float32 `thoth:"eq=24"`
	Float64        float64  `thoth:"eq=25"`
	Float64Pointer *float64 `thoth:"eq=26"`

	// complex is the set of all complex numbers with float32 real and imaginary parts.
	Complex64         complex64   `thoth:"eq=27"`
	Complex64Pointer  *complex64  `thoth:"eq=28"`
	Complex128        complex128  `thoth:"eq=29"`
	Complex128Pointer *complex128 `thoth:"eq=30"`
}
