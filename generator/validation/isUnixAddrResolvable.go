package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUnixAddrResolvableInput TODO
type IsUnixAddrResolvableInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUnixAddrResolvable TODO
func IsUnixAddrResolvable(_buffer io.StringWriter, input *IsUnixAddrResolvableInput) {
	condition, isLoop := isFunc("isUnixAddrResolvable", input.Field, input.Ref, "unix_addr")
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
