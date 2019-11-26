package models

// HostnameRFC952Validate TODO
type HostnameRFC952Validate struct {
	HostnameRFC952 string   `thoth:"hostname"`
	Pointer        *string  `thoth:"hostname"`
	Slice          []string `thoth:"hostname"`
}

// HostnameRFC1123Validate TODO
type HostnameRFC1123Validate struct {
	HostnameRFC1123 string   `thoth:"hostname_rfc1123"`
	Pointer         *string  `thoth:"hostname_rfc1123"`
	Slice           []string `thoth:"hostname_rfc1123"`
}
