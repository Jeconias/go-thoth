package models

// TypeSliceString TODO
type TypeSliceString struct {
	SliceString               []string   `thoth:"required"`
	PointerSliceString        *[]string  `thoth:"required"`
	SlicePointerString        []*string  `thoth:"required"`
	PointerSlicePointerString *[]*string `thoth:"required"`
}

// TypeSliceUint TODO
type TypeSliceUint struct {
	SliceUint               []uint   `thoth:"required"`
	PointerSliceUint        *[]uint  `thoth:"required"`
	SlicePointerUint        []*uint  `thoth:"required"`
	PointerSlicePointerUint *[]*uint `thoth:"required"`
}

// SliceInt64 TODO
// TODO: To check if it's slice pointer types
type SliceInt64 struct {
	A []int64 `thoth:"required"`
	// B *[]int64 `thoth:"required"`
	C []*int64 `thoth:"required"`
	// D *[]*int64 `thoth:"required"`
}
