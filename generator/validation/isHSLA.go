package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHSLAInput TODO
type IsHSLAInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHSLA TODO
func IsHSLA(_buffer io.StringWriter, input *IsHSLAInput) {
	rules.RenderCondition(
		_buffer,
		isRegex("hslaRegex", input.Field, input.Ref, "hsla"),
		input.Field,
		input.Tag,
	)
}
