package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsRGBInput TODO
type IsRGBInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsRGB TODO
func IsRGB(_buffer io.StringWriter, input *IsRGBInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("rgbRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
