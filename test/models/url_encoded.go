package models

// URLEncodedValidate TODO
type URLEncodedValidate struct {
	URLEncoded string   `thoth:"url_encoded"`
	Pointer    *string  `thoth:"url_encoded"`
	Slice      []string `thoth:"url_encoded"`
}
