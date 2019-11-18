package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHSLInput TODO
type IsHSLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHSL TODO
func IsHSL(_buffer io.StringWriter, input *IsHSLInput) {
	rules.RenderCondition(
		_buffer,
		isRegex("hslRegex", input.Field, input.Ref, "hsl"),
		input.Field,
		input.Tag,
	)
}
