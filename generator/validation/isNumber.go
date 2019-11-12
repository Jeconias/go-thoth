package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsNumberInput TODO
type IsNumberInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsNumber TODO
func IsNumber(_buffer io.StringWriter, input *IsNumberInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("numberRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
