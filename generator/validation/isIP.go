package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsIPInput TODO
type IsIPInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsIP TODO
func IsIP(_buffer io.StringWriter, input *IsIPInput, v ...string) {
	var version string
	if len(v) > 0 {
		version = v[0]
	}

	condition, isLoop := isFunc(fmt.Sprintf("isIP%s", version), input.Field, input.Ref, fmt.Sprintf("ip%s", version))
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
