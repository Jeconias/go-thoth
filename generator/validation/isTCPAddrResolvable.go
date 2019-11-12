package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsTCPAddrResolvableInput TODO
type IsTCPAddrResolvableInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsTCPAddrResolvable TODO
func IsTCPAddrResolvable(_buffer io.StringWriter, input *IsTCPAddrResolvableInput, v ...string) {
	var version string
	if len(v) > 0 {
		version = v[0]
	}
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isTCPAddrResolvable%s(%s)", version, input.Ref),
		input.Field,
		input.Tag,
	)
}
