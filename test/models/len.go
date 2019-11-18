package models

// LenStringValidate TODO
type LenStringValidate struct {
	String  string  `thoth:"len=5"`
	Pointer *string `thoth:"len=5"`
}

// LenChanValidate TODO
type LenChanValidate struct {
	ChanString        chan string  `thoth:"len=2"`
	ChanStringPointer chan *string `thoth:"len=2"`
	ChanUint          chan uint    `thoth:"len=2"`
	ChanUintPointer   chan *uint   `thoth:"len=2"`
}

// LenSliceStringValidate TODO
type LenSliceStringValidate struct {
	String              []string   `thoth:"len=5"`
	SlicePointer        []*string  `thoth:"len=5"`
	PointerSlice        *[]string  `thoth:"len=5"`
	PointerSlicePointer *[]*string `thoth:"len=5"`
}

// LenSliceIntValidate TODO
type LenSliceIntValidate struct {
	Int                 []int   `thoth:"len=5"`
	SlicePointer        []*int  `thoth:"len=5"`
	PointerSlice        *[]int  `thoth:"len=5"`
	PointerSlicePointer *[]*int `thoth:"len=5"`
}

// LenSliceFloat64Validate TODO
type LenSliceFloat64Validate struct {
	Float64             []float64   `thoth:"len=5"`
	SlicePointer        []*float64  `thoth:"len=5"`
	PointerSlice        *[]float64  `thoth:"len=5"`
	PointerSlicePointer *[]*float64 `thoth:"len=5"`
}
