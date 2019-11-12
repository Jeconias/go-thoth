package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUIDInput TODO
type IsUUIDInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID TODO
func IsUUID(_buffer io.StringWriter, input *IsUUIDInput) {
	rules.RenderCondition(
		_buffer,
		regexMatch("uUIDRegex", input.Ref),
		input.Field,
		input.Tag,
	)
}
