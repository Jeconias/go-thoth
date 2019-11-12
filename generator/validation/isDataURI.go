package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsDataURIInput TODO
type IsDataURIInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsDataURI TODO
func IsDataURI(_buffer io.StringWriter, input *IsDataURIInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isDataURI(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
