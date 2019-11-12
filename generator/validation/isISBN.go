package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsISBNInput TODO
type IsISBNInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsISBN TODO
func IsISBN(_buffer io.StringWriter, input *IsISBNInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isISBN10(%s) || isISBN13(%s)", input.Ref, input.Ref),
		input.Field,
		input.Tag,
	)
}
