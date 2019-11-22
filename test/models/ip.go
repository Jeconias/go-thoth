package models

// IPValidate TODO
type IPValidate struct {
	IP      string   `thoth:"ip"`
	Pointer *string  `thoth:"ip"`
	Slice   []string `thoth:"ip"`
}

// IPv4Validate TODO
type IPv4Validate struct {
	IPv4    string   `thoth:"ipv4"`
	Pointer *string  `thoth:"ipv4"`
	Slice   []string `thoth:"ipv4"`
}

// IPv6Validate TODO
type IPv6Validate struct {
	IPv6    string   `thoth:"ipv6"`
	Pointer *string  `thoth:"ipv6"`
	Slice   []string `thoth:"ipv6"`
}
