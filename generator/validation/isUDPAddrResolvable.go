package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUDPAddrResolvableInput TODO
type IsUDPAddrResolvableInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUDPAddrResolvable TODO
func IsUDPAddrResolvable(_buffer io.StringWriter, input *IsUDPAddrResolvableInput, v ...string) {
	var version string
	if len(v) > 0 {
		version = v[0]
	}
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isUDPAddrResolvable%s(%s)", version, input.Ref),
		input.Field,
		input.Tag,
	)
}
