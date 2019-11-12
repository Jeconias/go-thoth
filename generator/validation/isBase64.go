package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsBase64Input TODO
type IsBase64Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsBase64 TODO
func IsBase64(_buffer io.StringWriter, input *IsBase64Input) {
	rules.RenderCondition(
		_buffer,
		regexMatch("base64Regex", input.Ref),
		input.Field,
		input.Tag,
	)
}
