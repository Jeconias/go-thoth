package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsNumericInput TODO
type IsNumericInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsNumeric TODO
func IsNumeric(_buffer io.StringWriter, input *IsNumericInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("numericRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
