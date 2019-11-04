package validators

import "github.com/lab259/errors/v2"

var (
	errRequiredGeneric = errors.New("string could not be empty")
)

// IsSliceString is the validation if is empty
func IsSliceString(v []string) error {
	if len(v) == 0 {
		return errRequiredGeneric
	}
	return nil
}

// IsSliceStringPtr is the validation if is empty
func IsSliceStringPtr(v *[]string) error {
	if isValid(v) {
		return errRequiredGeneric
	}

	return IsSliceString(*v)
}

// IsSlicePtrString is the validation if is empty
func IsSlicePtrString(v []*string) error {
	if isValid(v) {
		return errRequiredGeneric
	}
	if len(v) == 0 {
		return errRequiredGeneric
	}
	return nil
}

// IsSlicePtrStringPtr is the validation if is empty
func IsSlicePtrStringPtr(v *[]*string) error {
	if isValid(v) {
		return errRequiredGeneric
	}

	return IsSlicePtrString(*v)
}
