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
	condition, isLoop := isRegex("uUID3Regex", input.Field, input.Ref, "uuid3")
	if isLoop {
		rules.RenderLoop(
			_buffer,
			condition,
			input.Ref,
			input.Field,
			input.Tag,
		)
	} else {
		rules.RenderCondition(
			_buffer,
			condition,
			input.Field,
			input.Tag,
		)
	}
}
