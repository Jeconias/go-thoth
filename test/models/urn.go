package models

// URNValidate TODO
type URNValidate struct {
	URN     string   `thoth:"urn_rfc2141"`
	Pointer *string  `thoth:"urn_rfc2141"`
	Slice   []string `thoth:"urn_rfc2141"`
}
