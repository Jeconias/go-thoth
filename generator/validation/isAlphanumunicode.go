package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsAlphanumunicodeInput TODO
type IsAlphanumunicodeInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsAlphanumunicode TODO
func IsAlphanumunicode(_buffer io.StringWriter, input *IsAlphanumunicodeInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("alphaUnicodeNumericRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
