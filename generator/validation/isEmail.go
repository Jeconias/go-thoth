package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsEmailInput TODO
type IsEmailInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsEmail TODO
func IsEmail(_buffer io.StringWriter, input *IsEmailInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("emailRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
