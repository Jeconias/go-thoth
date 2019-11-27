package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasMinOfInput TODO
type HasMinOfInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// HasMinOf TODO
func HasMinOf(_buffer io.StringWriter, input *HasMinOfInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isCompare(&ValidateInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
			Value: input.Value,
		}, "<=", "min"),
		input.Field,
		input.Tag,
	)
}
