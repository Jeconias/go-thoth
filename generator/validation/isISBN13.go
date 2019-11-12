package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsISBN13Input TODO
type IsISBN13Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsISBN13 TODO
func IsISBN13(_buffer io.StringWriter, input *IsISBN13Input) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isISBN13(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
