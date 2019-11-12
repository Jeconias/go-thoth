package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasMaxOfInput TODO
type HasMaxOfInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// HasMaxOf TODO
func HasMaxOf(_buffer io.StringWriter, input *HasMaxOfInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		hasMaxOf(input),
		input.Field,
		input.Tag,
	)
}

func hasMaxOf(input *HasMaxOfInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return " len(" + input.Ref + ") < " + input.Value.(string)
		case *myasthurts.StarRefType:
			return input.Ref + " == nil"
		}
	}
	return ""
}
