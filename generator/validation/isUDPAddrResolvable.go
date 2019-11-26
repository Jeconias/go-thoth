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

	condition, isLoop := isFunc(fmt.Sprintf("isUDP%sAddrResolvable", version), input.Field, input.Ref, fmt.Sprintf("udp%s_addr", version))
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
