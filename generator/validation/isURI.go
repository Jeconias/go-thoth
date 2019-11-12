package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURIInput TODO
type IsURIInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURI TODO
func IsURI(_buffer io.StringWriter, input *IsURIInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isURI(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
