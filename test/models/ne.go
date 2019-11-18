package models

// NeNumber TODO
type NeNumber struct {
	//  uint is an unsigned integer type
	Uint           uint     `thoth:"ne=1"`
	UintPointer    *uint    `thoth:"ne=2"`
	Uint8          uint8    `thoth:"ne=3"`
	Uint8Pointer   *uint8   `thoth:"ne=4"`
	Uint16         uint16   `thoth:"ne=5"`
	Uint16Pointer  *uint16  `thoth:"ne=6"`
	Uint32         uint32   `thoth:"ne=7"`
	Uint32Pointer  *uint32  `thoth:"ne=8"`
	Uint64         uint64   `thoth:"ne=9"`
	Uint64Pointer  *uint64  `thoth:"ne=10"`
	Uintptr        uintptr  `thoth:"ne=11"`
	UintptrPointer *uintptr `thoth:"ne=12"`

	// int is a signed integer type
	Int          int    `thoth:"ne=13"`
	IntPointer   *int   `thoth:"ne=14"`
	Int8         int8   `thoth:"ne=15"`
	Int8Pointer  *int8  `thoth:"ne=16"`
	Int16        int16  `thoth:"ne=17"`
	Int16Pointer *int16 `thoth:"ne=18"`
	Int32        int32  `thoth:"ne=19"`
	Int32Pointer *int32 `thoth:"ne=20"`
	Int64        int64  `thoth:"ne=21"`
	Int64Pointer *int64 `thoth:"ne=22"`

	// float is the set of all IEEE-754
	Float32        float32  `thoth:"ne=23"`
	Float32Pointer *float32 `thoth:"ne=24"`
	Float64        float64  `thoth:"ne=25"`
	Float64Pointer *float64 `thoth:"ne=26"`

	// complex is the set of all complex numbers with float32 real and imaginary parts.
	Complex64         complex64   `thoth:"ne=27"`
	Complex64Pointer  *complex64  `thoth:"ne=28"`
	Complex128        complex128  `thoth:"ne=29"`
	Complex128Pointer *complex128 `thoth:"ne=30"`
}

// NeString TODO
type NeString struct {
	String  string  `thoth:"ne=chico"`
	Pointer *string `thoth:"ne=bento"`
}
