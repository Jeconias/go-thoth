package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsAlphaInput TODO
type IsAlphaInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsAlpha TODO
func IsAlpha(_buffer io.StringWriter, input *IsAlphaInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("alphaRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
