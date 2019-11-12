package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsIPAddrResolvableInput TODO
type IsIPAddrResolvableInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsIPAddrResolvable TODO
func IsIPAddrResolvable(_buffer io.StringWriter, input *IsIPAddrResolvableInput, v ...string) {
	var version string
	if len(v) > 0 {
		version = v[0]
	}
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isIPAddrResolvable%s(%s)", version, input.Ref),
		input.Field,
		input.Tag,
	)
}
