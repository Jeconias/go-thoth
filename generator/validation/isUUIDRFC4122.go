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
	rules.RenderCondition(
		_buffer,
		regexMatch("uUID5RFC4122Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
