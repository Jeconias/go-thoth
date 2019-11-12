package validation

import (
	"fmt"
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
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isUnixAddrResolvable(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
