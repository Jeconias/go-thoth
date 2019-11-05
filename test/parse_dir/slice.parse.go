package any

// SliceString TODO
type SliceString struct {
	A []string   `thoth:"required"`
	B *[]string  `thoth:"required"`
	C []*string  `thoth:"required"`
	D *[]*string `thoth:"required"`
}

// SliceInt64 TODO
// TODO: To check if it's slice pointer types
type SliceInt64 struct {
	A []int64 `thoth:"required"`
	// B *[]int64 `thoth:"required"`
	C []*int64 `thoth:"required"`
	// D *[]*int64 `thoth:"required"`
}
