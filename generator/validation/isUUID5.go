package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUUID5Input TODO
type IsUUID5Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUUID5 TODO
func IsUUID5(_buffer io.StringWriter, input *IsUUID5Input) {
	rules.RenderCondition(
		_buffer,
		regexMatch("uUID5Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
