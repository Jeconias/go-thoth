package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsSSNInput TODO
type IsSSNInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsSSN TODO
func IsSSN(_buffer io.StringWriter, input *IsSSNInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isSSN(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
