package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsFileInput TODO
type IsFileInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsFile TODO
func IsFile(_buffer io.StringWriter, input *IsFileInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isFile(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
