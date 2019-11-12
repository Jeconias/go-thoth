package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHostnameRFC952Input TODO
type IsHostnameRFC952Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHostnameRFC952 TODO
func IsHostnameRFC952(_buffer io.StringWriter, input *IsHostnameRFC952Input) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isHostnameRFC952(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
