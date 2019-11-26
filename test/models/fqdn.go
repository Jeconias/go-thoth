package models

// FQDNValidate TODO
type FQDNValidate struct {
	FQDN    string   `thoth:"fqdn"`
	Pointer *string  `thoth:"fqdn"`
	Slice   []string `thoth:"fqdn"`
}
