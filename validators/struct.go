package validators

import (
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

// Level wrap from validator.
type Level struct {
	*validator.Validate
}

// Top returns the top level struct
func (v *Level) Top() reflect.Value {
	return v.Top()
	// panic("not implemented")
}

// Parent returns the current struct parent
func (v *Level) Parent() reflect.Value {
	return v.Parent()
	// panic("not implemented")
}

// Current returns the current struct.
func (v *Level) Current() reflect.Value {
	return v.Current()
	// panic("not implemented")
}

// Validator returns the main validation object, in case one want to call validations internally.
func (v *Level) Validator() *validator.Validate {
	return v.Validate
	// panic("not implemented")
}

// ExtractType gets the actual underlying type of field value.
func (v *Level) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	return v.ExtractType(field)
	// panic("not implemented")
}

// ReportError reports an error just by passing the field and tag information
func (v *Level) ReportError(field interface{}, fieldName, structFieldName, tag, param string) {
	v.ReportError(field, fieldName, structFieldName, tag, param)
	// panic("not implemented")
}

// ReportValidationErrors reports ValidationErrors obtained from running validations within the Struct Level validation.
func (v *Level) ReportValidationErrors(relativeNamespace, relativeStructNamespace string, errs validator.ValidationErrors) {
	v.ReportValidationErrors(relativeNamespace, relativeStructNamespace, errs)
	// panic("not implemented")
}
