package models

// URLValidate TODO
// TODO: To improve pointer
type URLValidate struct {
	URL     string   `thoth:"url"`
	Pointer *string  `thoth:"url"`
	Slice   []string `thoth:"url"`
	// SlicePointer        []*string  `thoth:"url"`
	// PointerSlice        *[]string  `thoth:"url"`
	// PointerSlicePointer *[]*string `thoth:"url"`
}
