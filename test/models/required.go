package models

// RequiredChanString TODO
type RequiredChanString struct {
	ChanString             chan string    `thoth:"required"`
	PointerChanString      chan *string   `thoth:"required"`
	ChanSliceString        chan []string  `thoth:"required"`
	ChanSlicePointerString chan []*string `thoth:"required"`
}

// RequiredChanUint TODO
type RequiredChanUint struct {
	ChanUint             chan uint    `thoth:"required"`
	PointerChanUint      chan *uint   `thoth:"required"`
	ChanSliceUint        chan []uint  `thoth:"required"`
	ChanSlicePointerUint chan []*uint `thoth:"required"`
}

// RequiredString TODO
type RequiredString struct {
	String  string  `thoth:"required"`
	Pointer *string `thoth:"required"`
}

// RequiredInterface TODO
type RequiredInterface struct {
	Interface        interface{}  `thoth:"required"`
	PointerInterface *interface{} `thoth:"required"`
}

// RequiredStructA TODO
type RequiredStructA struct {
	A string
	B int
}

// RequiredStruct TODO
type RequiredStruct struct {
	Struct                    RequiredStructA     `thoth:"required"`
	StructPointer             *RequiredStructA    `thoth:"required"`
	SliceStruct               []RequiredStructA   `thoth:"required"`
	SliceStructPointer        []*RequiredStructA  `thoth:"required"`
	SlicePointerStruct        *[]RequiredStructA  `thoth:"required"`
	SlicePointerStructPointer *[]*RequiredStructA `thoth:"required"`
}

// RequiredSliceString TODO
type RequiredSliceString struct {
	SliceString               []string   `thoth:"required"`
	PointerSliceString        *[]string  `thoth:"required"`
	SlicePointerString        []*string  `thoth:"required"`
	PointerSlicePointerString *[]*string `thoth:"required"`
}

/* Numbers: uint */

// RequiredSliceUint TODO
type RequiredSliceUint struct {
	SliceUint               []uint   `thoth:"required"`
	PointerSliceUint        *[]uint  `thoth:"required"`
	SlicePointerUint        []*uint  `thoth:"required"`
	PointerSlicePointerUint *[]*uint `thoth:"required"`
}

// RequiredSliceUint8 TODO
type RequiredSliceUint8 struct {
	SliceUint8               []uint8   `thoth:"required"`
	PointerSliceUint8        *[]uint8  `thoth:"required"`
	SlicePointerUint8        []*uint8  `thoth:"required"`
	PointerSlicePointerUint8 *[]*uint8 `thoth:"required"`
}

// RequiredSliceUint16 TODO
type RequiredSliceUint16 struct {
	SliceUint16               []uint16   `thoth:"required"`
	PointerSliceUint16        *[]uint16  `thoth:"required"`
	SlicePointerUint16        []*uint16  `thoth:"required"`
	PointerSlicePointerUint16 *[]*uint16 `thoth:"required"`
}

// RequiredSliceUint32 TODO
type RequiredSliceUint32 struct {
	SliceUint32               []uint32   `thoth:"required"`
	PointerSliceUint32        *[]uint32  `thoth:"required"`
	SlicePointerUint32        []*uint32  `thoth:"required"`
	PointerSlicePointerUint32 *[]*uint32 `thoth:"required"`
}

// RequiredSliceUint64 TODO
type RequiredSliceUint64 struct {
	SliceUint64               []uint64   `thoth:"required"`
	PointerSliceUint64        *[]uint64  `thoth:"required"`
	SlicePointerUint64        []*uint64  `thoth:"required"`
	PointerSlicePointerUint64 *[]*uint64 `thoth:"required"`
}

// RequiredSliceUintptr TODO
type RequiredSliceUintptr struct {
	SliceUintptr               []uintptr   `thoth:"required"`
	PointerSliceUintptr        *[]uintptr  `thoth:"required"`
	SlicePointerUintptr        []*uintptr  `thoth:"required"`
	PointerSlicePointerUintptr *[]*uintptr `thoth:"required"`
}

/* Numbers: int */

// RequiredSliceInt TODO
type RequiredSliceInt struct {
	SliceInt               []int   `thoth:"required"`
	PointerSliceInt        *[]int  `thoth:"required"`
	SlicePointerInt        []*int  `thoth:"required"`
	PointerSlicePointerInt *[]*int `thoth:"required"`
}

// RequiredSliceInt8 TODO
type RequiredSliceInt8 struct {
	SliceInt8               []int8   `thoth:"required"`
	PointerSliceInt8        *[]int8  `thoth:"required"`
	SlicePointerInt8        []*int8  `thoth:"required"`
	PointerSlicePointerInt8 *[]*int8 `thoth:"required"`
}

// RequiredSliceInt16 TODO
type RequiredSliceInt16 struct {
	SliceInt16               []int16   `thoth:"required"`
	PointerSliceInt16        *[]int16  `thoth:"required"`
	SlicePointerInt16        []*int16  `thoth:"required"`
	PointerSlicePointerInt16 *[]*int16 `thoth:"required"`
}

// RequiredSliceInt32 TODO
type RequiredSliceInt32 struct {
	SliceInt32               []int32   `thoth:"required"`
	PointerSliceInt32        *[]int32  `thoth:"required"`
	SlicePointerInt32        []*int32  `thoth:"required"`
	PointerSlicePointerInt32 *[]*int32 `thoth:"required"`
}

// RequiredSliceInt64 TODO
type RequiredSliceInt64 struct {
	SliceInt64               []int64   `thoth:"required"`
	PointerSliceInt64        *[]int64  `thoth:"required"`
	SlicePointerInt64        []*int64  `thoth:"required"`
	PointerSlicePointerInt64 *[]*int64 `thoth:"required"`
}

/* Numbers: float */

// RequiredSliceFloat32 TODO
type RequiredSliceFloat32 struct {
	SliceFloat32               []float32   `thoth:"required"`
	PointerSliceFloat32        *[]float32  `thoth:"required"`
	SlicePointerFloat32        []*float32  `thoth:"required"`
	PointerSlicePointerFloat32 *[]*float32 `thoth:"required"`
}

// RequiredSliceFloat64 TODO
type RequiredSliceFloat64 struct {
	SliceFloat64               []float64   `thoth:"required"`
	PointerSliceFloat64        *[]float64  `thoth:"required"`
	SlicePointerFloat64        []*float64  `thoth:"required"`
	PointerSlicePointerFloat64 *[]*float64 `thoth:"required"`
}

/* Numbers: complex */

// RequiredSliceComplex64 TODO
type RequiredSliceComplex64 struct {
	SliceComplex64               []complex64   `thoth:"required"`
	PointerSliceComplex64        *[]complex64  `thoth:"required"`
	SlicePointerComplex64        []*complex64  `thoth:"required"`
	PointerSlicePointerComplex64 *[]*complex64 `thoth:"required"`
}

// RequiredSliceComplex128 TODO
type RequiredSliceComplex128 struct {
	SliceComplex128               []complex128   `thoth:"required"`
	PointerSliceComplex128        *[]complex128  `thoth:"required"`
	SlicePointerComplex128        []*complex128  `thoth:"required"`
	PointerSlicePointerComplex128 *[]*complex128 `thoth:"required"`
}

// RequiredNumber TODO
type RequiredNumber struct {
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

// RequiredMapStringToInterface TODO
type RequiredMapStringToInterface struct {
	MapStringToInterface        map[string]interface{}  `thoth:"required"`
	PointerMapStringToInterface *map[string]interface{} `thoth:"required"`
}

// RequiredMapStringToString TODO
type RequiredMapStringToString struct {
	MapStringToString        map[string]string  `thoth:"required"`
	PointerMapStringToString *map[string]string `thoth:"required"`
}

// RequiredMapIntToString TODO
type RequiredMapIntToString struct {
	MapIntToString        map[int]string  `thoth:"required"`
	PointerMapIntToString *map[int]string `thoth:"required"`
}

// RequiredMapIntToBool TODO
type RequiredMapIntToBool struct {
	MapIntToBool        map[int]bool  `thoth:"required"`
	PointerMapIntToBool *map[int]bool `thoth:"required"`
}

// MapTypeA TODO
type MapTypeA struct {
	Int  int
	Bool bool
}

// MapTypeB TODO
type MapTypeB struct {
	Bool   bool
	String string
	Float  float64
}

// RequiredMapIntToStruct TODO
type RequiredMapIntToStruct struct {
	MapIntToStructA               map[int]MapTypeA   `thoth:"required"`
	PointerMapIntToStructA        map[int]*MapTypeA  `thoth:"required"`
	MapIntToStructB               *map[int]MapTypeB  `thoth:"required"`
	PointerMapIntToStructPointerB *map[int]*MapTypeB `thoth:"required"`
}
