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
	rules.RenderCondition(
		_buffer,
		regexMatch("uUID5RFC4122Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
