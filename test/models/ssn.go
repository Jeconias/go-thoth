package models

// SSNValidate TODO
type SSNValidate struct {
	SSN     string   `thoth:"ssn"`
	Pointer *string  `thoth:"ssn"`
	Slice   []string `thoth:"ssn"`
}
