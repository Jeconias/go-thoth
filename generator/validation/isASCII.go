package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsASCIIInput TODO
type IsASCIIInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsASCII TODO
func IsASCII(_buffer io.StringWriter, input *IsASCIIInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("aSCIIRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
