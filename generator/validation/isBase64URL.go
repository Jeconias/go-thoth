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
	rules.RenderCondition(
		_buffer,
		regexMatch("base64URLRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
