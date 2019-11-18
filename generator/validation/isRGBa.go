package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsRGBaInput TODO
type IsRGBaInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsRGBa TODO
func IsRGBa(_buffer io.StringWriter, input *IsRGBaInput) {
	rules.RenderCondition(
		_buffer,
		isRegex("rgbaRegex", input.Field, input.Ref, "rgba"),
		input.Field,
		input.Tag,
	)
}
