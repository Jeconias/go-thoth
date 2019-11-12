package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHTMLInput TODO
type IsHTMLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHTML TODO
func IsHTML(_buffer io.StringWriter, input *IsHTMLInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isHTML(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
