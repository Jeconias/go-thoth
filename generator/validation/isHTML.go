package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHTMLInput TODO
type IsHTMLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHTML TODO
func IsHTML(_buffer io.StringWriter, input *IsHTMLInput) {
	condition, isLoop := isRegex("hTMLRegex", input.Field, input.Ref, "html")
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
