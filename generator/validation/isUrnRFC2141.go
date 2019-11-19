package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUrnRFC2141Input TODO
type IsUrnRFC2141Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUrnRFC2141 TODO
func IsUrnRFC2141(_buffer io.StringWriter, input *IsUrnRFC2141Input) {
	condition, isLoop := isFunc("isUrnRFC2141", input.Field, input.Ref, "urn_rfc2141")
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
