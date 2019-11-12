package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsPrintableASCIIInput TODO
type IsPrintableASCIIInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsPrintableASCII TODO
func IsPrintableASCII(_buffer io.StringWriter, input *IsPrintableASCIIInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("printableASCIIRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
