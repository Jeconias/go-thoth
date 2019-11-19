package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUID5RFC4122Input TODO
type IsUUID5RFC4122Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID5RFC4122 TODO
func IsUUID5RFC4122(_buffer io.StringWriter, input *IsUUID5RFC4122Input) {
	condition, isLoop := isRegex("uUID5RFC4122Regex", input.Field, input.Ref, "uuid5_rfc4122")
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
