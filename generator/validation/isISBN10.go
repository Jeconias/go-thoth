package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsISBN10Input TODO
type IsISBN10Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsISBN10 TODO
func IsISBN10(_buffer io.StringWriter, input *IsISBN10Input) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isISBN10(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
