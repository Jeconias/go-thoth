package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUID4Input TODO
type IsUUID4Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID4 TODO
func IsUUID4(_buffer io.StringWriter, input *IsUUID4Input) {
	rules.RenderCondition(
		_buffer,
		regexMatch("uUID4Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
