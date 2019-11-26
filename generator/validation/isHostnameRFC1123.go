package validation

import (
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
	condition, isLoop := isRegex("hostnameRegexRFC1123", input.Field, input.Ref, "hostname")
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
