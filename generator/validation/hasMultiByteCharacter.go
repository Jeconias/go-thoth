package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasMultiByteCharacterInput TODO
type HasMultiByteCharacterInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// HasMultiByteCharacter TODO
func HasMultiByteCharacter(_buffer io.StringWriter, input *HasMultiByteCharacterInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("multibyteRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
