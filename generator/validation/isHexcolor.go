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
	rules.RenderCondition(
		_buffer,
		isRegex("hexcolorRegex", input.Field, input.Ref, "hexcolor"),
		input.Field,
		input.Tag,
	)
}
