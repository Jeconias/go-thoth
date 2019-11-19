package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUIDRFC4122Input TODO
type IsUUIDRFC4122Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUIDRFC4122 TODO
func IsUUIDRFC4122(_buffer io.StringWriter, input *IsUUIDRFC4122Input) {
	condition, isLoop := isRegex("uUIDRFC4122Regex", input.Field, input.Ref, "uuid_rfc4122")
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
