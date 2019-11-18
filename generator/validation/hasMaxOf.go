package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasMaxOfInput TODO
type HasMaxOfInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// HasMaxOf TODO
func HasMaxOf(_buffer io.StringWriter, input *HasMaxOfInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isCompare(&IsCompareInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
			Value: input.Value,
		}, ">=", "max"),
		input.Field,
		input.Tag,
	)
}
