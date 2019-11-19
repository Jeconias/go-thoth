package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsBase64URLInput TODO
type IsBase64URLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsBase64URL TODO
func IsBase64URL(_buffer io.StringWriter, input *IsBase64URLInput) {
	condition, isLoop := isRegex("base64URLRegex", input.Field, input.Ref, "base64url")
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
