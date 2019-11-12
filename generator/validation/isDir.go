package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsDirInput TODO
type IsDirInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsDir TODO
func IsDir(_buffer io.StringWriter, input *IsDirInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isDir(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
