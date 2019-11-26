package models

// UDPAddrValidate TODO
type UDPAddrValidate struct {
	UDPAddr string   `thoth:"udp_addr"`
	Pointer *string  `thoth:"udp_addr"`
	Slice   []string `thoth:"udp_addr"`
}

// UDP4AddrValidate TODO
type UDP4AddrValidate struct {
	UDP4Addr string   `thoth:"udp4_addr"`
	Pointer  *string  `thoth:"udp4_addr"`
	Slice    []string `thoth:"udp4_addr"`
}

// UDP6AddrValidate TODO
type UDP6AddrValidate struct {
	UDP6Addr string   `thoth:"udp6_addr"`
	Pointer  *string  `thoth:"udp6_addr"`
	Slice    []string `thoth:"udp6_addr"`
}
