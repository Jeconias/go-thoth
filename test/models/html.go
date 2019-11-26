package models

// HTMLValidate TODO
type HTMLValidate struct {
	HTML    string   `thoth:"html"`
	Pointer *string  `thoth:"html"`
	Slice   []string `thoth:"html"`
}

// HTMLEncodedValidate TODO
type HTMLEncodedValidate struct {
	HTMLEncoded string   `thoth:"html_encoded"`
	Pointer     *string  `thoth:"html_encoded"`
	Slice       []string `thoth:"html_encoded"`
}
