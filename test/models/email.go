package models

// EmailValidate TODO
type EmailValidate struct {
	Email   string  `thoth:"email"`
	Pointer *string `thoth:"email"`
}

// EmailSliceValidate TODO
type EmailSliceValidate struct {
	Emails []string `thoth:"email"`
	// SlicePointer        []*string `thoth:"email"`
	// PointerSlice        *[]string  `thoth:"email"`
	// PointerSlicePointer *[]*string `thoth:"email"`
}
