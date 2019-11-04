package validators

import (
	"github.com/lab259/errors/v2"
)

var (
	errRequiredString = errors.New("string could not be empty")
)

// Empty is the validation if is empty
func Empty(v string) error {
	if v == "" {
		return errRequiredString
	}
	return nil
}

// EmptyPtr is the validation if is empty
func EmptyPtr(v *string) error {
	if isValid(v) {
		return errRequiredString
	}

	return Empty(*v)
}
