package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHexadecimalInput TODO
type IsHexadecimalInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHexadecimal TODO
func IsHexadecimal(_buffer io.StringWriter, input *IsHexadecimalInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("hexadecimalRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
