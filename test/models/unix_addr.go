package models

// UNIXAddrValidate TODO
type UNIXAddrValidate struct {
	UNIXAddr string   `thoth:"unix_addr"`
	Pointer  *string  `thoth:"unix_addr"`
	Slice    []string `thoth:"unix_addr"`
}
