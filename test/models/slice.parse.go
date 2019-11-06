package models

// TypeSliceString TODO
type TypeSliceString struct {
	SliceString               []string   `thoth:"required"`
	PointerSliceString        *[]string  `thoth:"required"`
	SlicePointerString        []*string  `thoth:"required"`
	PointerSlicePointerString *[]*string `thoth:"required"`
}

/* Numbers: uint */

// TypeSliceUint TODO
type TypeSliceUint struct {
	SliceUint               []uint   `thoth:"required"`
	PointerSliceUint        *[]uint  `thoth:"required"`
	SlicePointerUint        []*uint  `thoth:"required"`
	PointerSlicePointerUint *[]*uint `thoth:"required"`
}

// TypeSliceUint8 TODO
type TypeSliceUint8 struct {
	SliceUint8               []uint8   `thoth:"required"`
	PointerSliceUint8        *[]uint8  `thoth:"required"`
	SlicePointerUint8        []*uint8  `thoth:"required"`
	PointerSlicePointerUint8 *[]*uint8 `thoth:"required"`
}

// TypeSliceUint16 TODO
type TypeSliceUint16 struct {
	SliceUint16               []uint16   `thoth:"required"`
	PointerSliceUint16        *[]uint16  `thoth:"required"`
	SlicePointerUint16        []*uint16  `thoth:"required"`
	PointerSlicePointerUint16 *[]*uint16 `thoth:"required"`
}

// TypeSliceUint32 TODO
type TypeSliceUint32 struct {
	SliceUint32               []uint32   `thoth:"required"`
	PointerSliceUint32        *[]uint32  `thoth:"required"`
	SlicePointerUint32        []*uint32  `thoth:"required"`
	PointerSlicePointerUint32 *[]*uint32 `thoth:"required"`
}

// TypeSliceUint64 TODO
type TypeSliceUint64 struct {
	SliceUint64               []uint64   `thoth:"required"`
	PointerSliceUint64        *[]uint64  `thoth:"required"`
	SlicePointerUint64        []*uint64  `thoth:"required"`
	PointerSlicePointerUint64 *[]*uint64 `thoth:"required"`
}

// TypeSliceUintptr TODO
type TypeSliceUintptr struct {
	SliceUintptr               []uintptr   `thoth:"required"`
	PointerSliceUintptr        *[]uintptr  `thoth:"required"`
	SlicePointerUintptr        []*uintptr  `thoth:"required"`
	PointerSlicePointerUintptr *[]*uintptr `thoth:"required"`
}

/* Numbers: int */

// TypeSliceInt TODO
type TypeSliceInt struct {
	SliceInt               []int   `thoth:"required"`
	PointerSliceInt        *[]int  `thoth:"required"`
	SlicePointerInt        []*int  `thoth:"required"`
	PointerSlicePointerInt *[]*int `thoth:"required"`
}

// TypeSliceInt8 TODO
type TypeSliceInt8 struct {
	SliceInt8               []int8   `thoth:"required"`
	PointerSliceInt8        *[]int8  `thoth:"required"`
	SlicePointerInt8        []*int8  `thoth:"required"`
	PointerSlicePointerInt8 *[]*int8 `thoth:"required"`
}

// TypeSliceInt16 TODO
type TypeSliceInt16 struct {
	SliceInt16               []int16   `thoth:"required"`
	PointerSliceInt16        *[]int16  `thoth:"required"`
	SlicePointerInt16        []*int16  `thoth:"required"`
	PointerSlicePointerInt16 *[]*int16 `thoth:"required"`
}

// TypeSliceInt32 TODO
type TypeSliceInt32 struct {
	SliceInt32               []int32   `thoth:"required"`
	PointerSliceInt32        *[]int32  `thoth:"required"`
	SlicePointerInt32        []*int32  `thoth:"required"`
	PointerSlicePointerInt32 *[]*int32 `thoth:"required"`
}

// TypeSliceInt64 TODO
type TypeSliceInt64 struct {
	SliceInt64               []int64   `thoth:"required"`
	PointerSliceInt64        *[]int64  `thoth:"required"`
	SlicePointerInt64        []*int64  `thoth:"required"`
	PointerSlicePointerInt64 *[]*int64 `thoth:"required"`
}

/* Numbers: float */

// TypeSliceFloat32 TODO
type TypeSliceFloat32 struct {
	SliceFloat32               []float32   `thoth:"required"`
	PointerSliceFloat32        *[]float32  `thoth:"required"`
	SlicePointerFloat32        []*float32  `thoth:"required"`
	PointerSlicePointerFloat32 *[]*float32 `thoth:"required"`
}

// TypeSliceFloat64 TODO
type TypeSliceFloat64 struct {
	SliceFloat64               []float64   `thoth:"required"`
	PointerSliceFloat64        *[]float64  `thoth:"required"`
	SlicePointerFloat64        []*float64  `thoth:"required"`
	PointerSlicePointerFloat64 *[]*float64 `thoth:"required"`
}

/* Numbers: complex */

// TypeSliceComplex64 TODO
type TypeSliceComplex64 struct {
	SliceComplex64               []complex64   `thoth:"required"`
	PointerSliceComplex64        *[]complex64  `thoth:"required"`
	SlicePointerComplex64        []*complex64  `thoth:"required"`
	PointerSlicePointerComplex64 *[]*complex64 `thoth:"required"`
}

// TypeSliceComplex128 TODO
type TypeSliceComplex128 struct {
	SliceComplex128               []complex128   `thoth:"required"`
	PointerSliceComplex128        *[]complex128  `thoth:"required"`
	SlicePointerComplex128        []*complex128  `thoth:"required"`
	PointerSlicePointerComplex128 *[]*complex128 `thoth:"required"`
}
