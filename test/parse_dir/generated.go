//go:generate gofmt -s -w generated.go
package any

// Thoth validate
func (h *Home) Thoth() error {
	if required("int64", h.ID) {
		return errRequiredString
	}
	if required("string", h.Address) {
		return errRequiredString
	}
	return nil
}

// Thoth validate
func (u *User) Thoth() error {
	if required("string", u.Name) {
		return errRequiredString
	}
	return nil
}
