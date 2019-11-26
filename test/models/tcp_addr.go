package models

// TCPAddrValidate TODO
type TCPAddrValidate struct {
	TCPAddr string   `thoth:"tcp_addr"`
	Pointer *string  `thoth:"tcp_addr"`
	Slice   []string `thoth:"tcp_addr"`
}

// TCP4AddrValidate TODO
type TCP4AddrValidate struct {
	TCP4Addr string   `thoth:"tcp4_addr"`
	Pointer  *string  `thoth:"tcp4_addr"`
	Slice    []string `thoth:"tcp4_addr"`
}

// TCP6AddrValidate TODO
type TCP6AddrValidate struct {
	TCP6Addr string   `thoth:"tcp6_addr"`
	Pointer  *string  `thoth:"tcp6_addr"`
	Slice    []string `thoth:"tcp6_addr"`
}
