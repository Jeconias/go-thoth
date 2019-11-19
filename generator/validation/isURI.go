package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURIInput TODO
type IsURIInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURI TODO
func IsURI(_buffer io.StringWriter, input *IsURIInput) {
	condition, isLoop := isFunc("isURI", input.Field, input.Ref, "url")
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
