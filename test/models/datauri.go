package models

// DataURIValidate TODO
type DataURIValidate struct {
	DataURI string   `thoth:"datauri"`
	Pointer *string  `thoth:"datauri"`
	Slice   []string `thoth:"datauri"`
}
