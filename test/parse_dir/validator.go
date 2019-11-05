package any

import "fmt"

const fieldErrMsg = "Error:Field validation for '%s' failed on the '%s' tag"

// FieldError contains all functions to get error details
type FieldError interface {
	Tag() string
	Field() string
	Value() interface{}
	Type() string
}

// ValidationErrors is an array of FieldError's for use in custom error
// messages post validation.
type ValidationErrors []FieldError

type fieldError struct {
	tag   string
	field string
	typ   string
	value interface{}
}

// NewFieldError TODO
func NewFieldError(field, typ, tag string) *fieldError {
	return &fieldError{
		tag:   tag,
		field: field,
		typ:   typ,
	}
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

// Value returns the actual fields value in case needed for creating the error
// message
func (f *fieldError) Value() interface{} {
	return f.value
}

// Type returns the Field's string Type
func (f *fieldError) Type() string {
	return f.typ
}

// Error returns the fieldError's error message
func (f *fieldError) Error() string {
	return fmt.Sprintf(fieldErrMsg, f.Field(), f.Tag())
}

var (
	// ErrEmpty TODO
	ErrEmpty = func(field, tag string) *fieldError {
		return &fieldError{
			tag:   tag,
			field: field,
		}
	}

	// ErrNumberRequired TODO
	ErrNumberRequired = func(field, tag string) *fieldError {
		return &fieldError{
			tag:   tag,
			field: field,
		}
	}

	// ErrStringEmpty TODO
	ErrStringEmpty = func(field, tag string) *fieldError {
		return NewFieldError(field, "string", tag)
	}
	// ErrArrayEmpty TODO
	ErrArrayEmpty = func(field, tag string) *fieldError {
		return NewFieldError(field, "slice", tag)
	}
)

/*
	Validators
*/

// Empty TODO
func Empty(v int) bool {
	return v == 0
}

// Empty TODO
func IsValid(v interface{}) bool {
	return v == nil
}

/*
	Numbers
*/

// IsUint TODO
func IsUint(v *uint) bool {
	return v == nil || *v == 0
}

// IsUint8 TODO
func IsUint8(v *uint8) bool {
	return v == nil || *v == 0
}

// IsUint16 TODO
func IsUint16(v *uint16) bool {
	return v == nil || *v == 0
}

// IsUint32 TODO
func IsUint32(v *uint32) bool {
	return v == nil || *v == 0
}

// IsUint64 TODO
func IsUint64(v *uint64) bool {
	return v == nil || *v == 0
}

// IsUintptr TODO
func IsUintptr(v *uintptr) bool {
	return v == nil || *v == 0
}

// IsInt TODO
func IsInt(v *int) bool {
	return v == nil || *v == 0
}

// IsInt8 TODO
func IsInt8(v *int8) bool {
	return v == nil || *v == 0
}

// IsInt16 TODO
func IsInt16(v *int16) bool {
	return v == nil || *v == 0
}

// IsInt32 TODO
func IsInt32(v *int32) bool {
	return v == nil || *v == 0
}

// IsInt64 TODO
func IsInt64(v *int64) bool {
	return v == nil || *v == 0
}

// Is2Int64 TODO
func Is2Int64(v int) bool {
	return v == 0
}

// IsFloat32 TODO
func IsFloat32(v *float32) bool {
	return v == nil || *v == 0
}

// IsFloat64 TODO
func IsFloat64(v *float64) bool {
	return v == nil || *v == 0
}

// IsComplex64 TODO
func IsComplex64(v *complex64) bool {
	return v == nil || *v == 0
}

// IsComplex128 TODO
func IsComplex128(v *complex128) bool {
	return v == nil || *v == 0
}
