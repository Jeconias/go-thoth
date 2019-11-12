package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUID3Input TODO
type IsUUID3Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID3 TODO
func IsUUID3(_buffer io.StringWriter, input *IsUUID3Input) {
	rules.RenderCondition(
		_buffer,
		regexMatch("uUID3Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
