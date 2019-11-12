package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURLEncodedInput TODO
type IsURLEncodedInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURLEncoded TODO
func IsURLEncoded(_buffer io.StringWriter, input *IsURLEncodedInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isURLEncoded(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
