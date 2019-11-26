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
	condition, isLoop := isRegex("hexadecimalRegex", input.Field, input.Ref, "hexadecimal")
	if isLoop {
		rules.RenderLoop(
			_buffer,
			condition,
			input.Ref,
			input.Field,
			input.Tag,
		)
	} else {
		rules.RenderCondition(
			_buffer,
			condition,
			input.Field,
			input.Tag,
		)
	}
}
