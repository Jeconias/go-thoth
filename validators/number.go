package validators

import (
	"github.com/lab259/errors/v2"
)

var errRequiredInt = errors.New("int could not be zero")

// IsUint is the validation if not equal zero
func IsUint(v uint) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUintPtr is the validation if is valid
func IsUintPtr(v *uint) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUint(*v)
}

// IsUint8 is the validation if not equal zero
func IsUint8(v uint8) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUint8Ptr is the validation if is valid
func IsUint8Ptr(v *uint8) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUint8(*v)
}

// IsUint16 is the validation if not equal zero
func IsUint16(v uint16) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUint16Ptr is the validation if is valid
func IsUint16Ptr(v *uint16) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUint16(*v)
}

// IsUint32 is the validation if not equal zero
func IsUint32(v uint32) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUint32Ptr is the validation if is valid
func IsUint32Ptr(v *uint32) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUint32(*v)
}

// IsUint64 is the validation if not equal zero
func IsUint64(v uint64) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUint64Ptr is the validation if is valid
func IsUint64Ptr(v *uint64) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUint64(*v)
}

// IsUintptr is the validation if not equal zero
func IsUintptr(v uintptr) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsUint64ptrPtr is the validation if is valid
func IsUint64ptrPtr(v *uintptr) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsUintptr(*v)
}

// IsInt is the validation if not equal zero
func IsInt(v int) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsIntPtr is the validation if is valid
func IsIntPtr(v *int) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsInt(*v)
}

// IsInt8 is the validation if not equal zero
func IsInt8(v int8) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsInt8Ptr is the validation if is valid
func IsInt8Ptr(v *int8) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsInt8(*v)
}

// IsInt16 is the validation if not equal zero
func IsInt16(v int16) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsInt16Ptr is the validation if is valid
func IsInt16Ptr(v *int16) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsInt16(*v)
}

// IsInt32 is the validation if not equal zero
func IsInt32(v int32) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsInt32Ptr is the validation if is valid
func IsInt32Ptr(v *int32) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsInt32(*v)
}

// IsInt64 is the validation if not equal zero
func IsInt64(v int64) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsInt64Ptr is the validation if is valid
func IsInt64Ptr(v *int64) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsInt64(*v)
}

// IsFloat32 is the validation if not equal zero
func IsFloat32(v float32) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsFloat32Ptr is the validation if is valid
func IsFloat32Ptr(v *float32) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsFloat32(*v)
}

// IsFloat64 is the validation if not equal zero
func IsFloat64(v float64) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsFloat64Ptr is the validation if is valid
func IsFloat64Ptr(v *float64) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsFloat64(*v)
}

// IsComplex64 is the validation if not equal zero
func IsComplex64(v complex64) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsComplex64Ptr is the validation if is valid
func IsComplex64Ptr(v *complex64) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsComplex64(*v)
}

// IsComplex128 is the validation if not equal zero
func IsComplex128(v complex128) error {
	if v != 0 {
		return errRequiredInt
	}
	return nil
}

// IsComplex128Ptr is the validation if is valid
func IsComplex128Ptr(v *complex128) error {
	if isValid(v) {
		return errRequiredInt
	}

	return IsComplex128(*v)
}
