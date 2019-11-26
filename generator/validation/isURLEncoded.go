package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURLEncodedInput TODO
type IsURLEncodedInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURLEncoded TODO
func IsURLEncoded(_buffer io.StringWriter, input *IsURLEncodedInput) {
	condition, isLoop := isRegex("uRLEncodedRegex", input.Field, input.Ref, "url_encoded")
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
