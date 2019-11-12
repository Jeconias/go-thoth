package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsUrnRFC2141Input TODO
type IsUrnRFC2141Input struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsUrnRFC2141 TODO
func IsUrnRFC2141(_buffer io.StringWriter, input *IsUrnRFC2141Input) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isUrnRFC2141(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
