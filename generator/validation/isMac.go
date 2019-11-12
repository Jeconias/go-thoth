package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsMacInput TODO
type IsMacInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsMac TODO
func IsMac(_buffer io.StringWriter, input *IsMacInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isMac(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
