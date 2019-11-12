package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHTMLEncodedInput TODO
type IsHTMLEncodedInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHTMLEncoded TODO
func IsHTMLEncoded(_buffer io.StringWriter, input *IsHTMLEncodedInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isHTMLEncoded(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
