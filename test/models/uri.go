package models

// URIValidate TODO
// TODO: To improve pointer
type URIValidate struct {
	URI     string   `thoth:"uri"`
	Pointer *string  `thoth:"uri"`
	Slice   []string `thoth:"uri"`
}
