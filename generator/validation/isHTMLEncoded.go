package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHTMLEncodedInput TODO
type IsHTMLEncodedInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHTMLEncoded TODO
func IsHTMLEncoded(_buffer io.StringWriter, input *IsHTMLEncodedInput) {
	condition, isLoop := isRegex("hTMLEncodedRegex", input.Field, input.Ref, "html_encoded")
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
