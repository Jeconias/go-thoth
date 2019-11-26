package validation

import (
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
	condition, isLoop := isRegex("hostnameRegexRFC952", input.Field, input.Ref, "hostname")
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
