package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsURLInput TODO
type IsURLInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsURL TODO
func IsURL(_buffer io.StringWriter, input *IsURLInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isURL(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
