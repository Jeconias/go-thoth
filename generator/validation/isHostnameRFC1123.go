package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHostnameRFC1123Input TODO
type IsHostnameRFC1123Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHostnameRFC1123 TODO
func IsHostnameRFC1123(_buffer io.StringWriter, input *IsHostnameRFC1123Input) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isHostnameRFC1123(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
