package models

// IPAddrValidate TODO
type IPAddrValidate struct {
	IPAddr  string   `thoth:"ip_addr"`
	Pointer *string  `thoth:"ip_addr"`
	Slice   []string `thoth:"ip_addr"`
}

// IP4AddrValidate TODO
type IP4AddrValidate struct {
	IP4Addr string   `thoth:"ip4_addr"`
	Pointer *string  `thoth:"ip4_addr"`
	Slice   []string `thoth:"ip4_addr"`
}

// IP6AddrValidate TODO
type IP6AddrValidate struct {
	IP6Addr string   `thoth:"ip6_addr"`
	Pointer *string  `thoth:"ip6_addr"`
	Slice   []string `thoth:"ip6_addr"`
}
