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

	condition, isLoop := isFunc(fmt.Sprintf("isTCP%sAddrResolvable", version), input.Field, input.Ref, fmt.Sprintf("tcp%s_addr", version))
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
