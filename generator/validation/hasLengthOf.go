package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasLengthOfInput TODO
type HasLengthOfInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// HasLengthOf TODO
func HasLengthOf(_buffer io.StringWriter, input *HasLengthOfInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		hasLengthOf(input),
		input.Field,
		input.Tag,
	)
}

func hasLengthOf(input *HasLengthOfInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return " len(" + input.Ref + ") == " + input.Value.(string)
		case *myasthurts.StarRefType:
			return input.Ref + " == nil"
		}
	}
	return ""
}
