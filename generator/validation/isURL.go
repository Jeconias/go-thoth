package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURLInput TODO
type IsURLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURL TODO
func IsURL(_buffer io.StringWriter, input *IsURLInput) {
	condition, isLoop := isFunc("isURL", input.Field, input.Ref, "url")
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
