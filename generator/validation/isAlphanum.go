package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsAlphanumInput TODO
type IsAlphanumInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsAlphanum TODO
func IsAlphanum(_buffer io.StringWriter, input *IsAlphanumInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("alphaNumericRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
