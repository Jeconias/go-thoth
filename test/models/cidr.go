package models

// CIDRValidate TODO
type CIDRValidate struct {
	CIDR    string   `thoth:"cidr"`
	Pointer *string  `thoth:"cidr"`
	Slice   []string `thoth:"cidr"`
}

// CIDRv4Validate TODO
type CIDRv4Validate struct {
	CIDRv4  string   `thoth:"cidrv4"`
	Pointer *string  `thoth:"cidrv4"`
	Slice   []string `thoth:"cidrv4"`
}

// CIDRv6Validate TODO
type CIDRv6Validate struct {
	CIDRv6  string   `thoth:"cidrv6"`
	Pointer *string  `thoth:"cidrv6"`
	Slice   []string `thoth:"cidrv6"`
}
