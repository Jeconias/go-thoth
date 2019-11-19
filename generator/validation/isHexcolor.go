package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHexcolorInput TODO
type IsHexcolorInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHexcolor TODO
func IsHexcolor(_buffer io.StringWriter, input *IsHexcolorInput) {
	condition, isLoop := isRegex("hexcolorRegex", input.Field, input.Ref, "hexcolor")
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
