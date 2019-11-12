package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsAlphaunicodeInput TODO
type IsAlphaunicodeInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsAlphaunicode TODO
func IsAlphaunicode(_buffer io.StringWriter, input *IsAlphaunicodeInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("alphaUnicodeRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
