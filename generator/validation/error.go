package validation

import (
	"fmt"
	"reflect"

	myasthurts "github.com/lab259/go-my-ast-hurts"
)

// NewErrTypeNotSupported TODO
func NewErrTypeNotSupported(rule string, field *myasthurts.Field) error {
	return fmt.Errorf(`rule "%s" does not support data type: "%s" of "%s"`, rule, reflect.TypeOf(field.RefType).String(), field.RefType.Name())
}
