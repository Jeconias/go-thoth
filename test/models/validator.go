package models

import "fmt"

const fieldErrMsg = "Error: Validation of field '%s' failed on tag '%s'"

// FieldError contains all functions to get error details
type FieldError interface {
	Tag() string
	Field() string
}

// ValidationErrors is an array of FieldError's for use in custom error
// messages post validation.
type ValidationErrors []FieldError

type fieldError struct {
	tag   string
	field string
}

// Tag returns the validation tag that failed.
func (f *fieldError) Tag() string {
	return f.tag
}

// Field returns the fields name with the tag name taking precedence over the
// fields actual name.
func (f *fieldError) Field() string {
	return f.field
}

// Error returns the fieldError's error message
func (f *fieldError) Error() string {
	return fmt.Sprintf(fieldErrMsg, f.Field(), f.Tag())
}

var (
	// NewError TODO
	NewError = func(field, tag string) *fieldError {
		return &fieldError{
			tag:   tag,
			field: field,
		}
	}
)

/*
	Validators
*/

// Empty TODO
func Empty(v int) bool {
	return v == 0
}

// IsValid TODO
func IsValid(v interface{}) bool {
	return v == nil
}

/*
	Numbers
*/

// IsUint TODO
func IsUint(v uint) bool {
	return v == 0
}

// IsUint8 TODO
func IsUint8(v uint8) bool {
	return v == 0
}

// IsUint16 TODO
func IsUint16(v uint16) bool {
	return v == 0
}

// IsUint32 TODO
func IsUint32(v uint32) bool {
	return v == 0
}

// IsUint64 TODO
func IsUint64(v uint64) bool {
	return v == 0
}

// IsUintptr TODO
func IsUintptr(v uintptr) bool {
	return v == 0
}

// IsInt TODO
func IsInt(v int) bool {
	return v == 0
}

// IsInt8 TODO
func IsInt8(v int8) bool {
	return v == 0
}

// IsInt16 TODO
func IsInt16(v int16) bool {
	return v == 0
}

// IsInt32 TODO
func IsInt32(v int32) bool {
	return v == 0
}

// IsInt64 TODO
func IsInt64(v int64) bool {
	return v == 0
}

// IsFloat32 TODO
func IsFloat32(v float32) bool {
	return v == 0
}

// IsFloat64 TODO
func IsFloat64(v float64) bool {
	return v == 0
}

// IsComplex64 TODO
func IsComplex64(v complex64) bool {
	return v == 0
}

// IsComplex128 TODO
func IsComplex128(v complex128) bool {
	return v == 0
}
