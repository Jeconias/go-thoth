package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsEqInput TODO
type IsEqInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsEq TODO
func IsEq(_buffer io.StringWriter, input *IsEqInput) {
	rules.RenderEq(_buffer, &rules.EqInput{
		Field: input.Field,
		Tag:   input.Tag,
		Ref:   input.Ref,
		Value: input.Value,
	})
}
