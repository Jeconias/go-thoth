package models

// ISBNValidate TODO
type ISBNValidate struct {
	ISBN    string   `thoth:"isbn"`
	Pointer *string  `thoth:"isbn"`
	Slice   []string `thoth:"isbn"`
}

// ISBN10Validate TODO
type ISBN10Validate struct {
	ISBN10  string   `thoth:"isbn10"`
	Pointer *string  `thoth:"isbn10"`
	Slice   []string `thoth:"isbn10"`
}

// ISBN13Validate TODO
type ISBN13Validate struct {
	ISBN13  string   `thoth:"isbn13"`
	Pointer *string  `thoth:"isbn13"`
	Slice   []string `thoth:"isbn13"`
}
