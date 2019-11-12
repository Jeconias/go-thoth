package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// HasMinOfInput TODO
type HasMinOfInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// HasMinOf TODO
func HasMinOf(_buffer io.StringWriter, input *HasMinOfInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		hasMinOf(input),
		input.Field,
		input.Tag,
	)
}

func hasMinOf(input *HasMinOfInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return " len(" + input.Ref + ") > " + input.Value.(string)
		case *myasthurts.StarRefType:
			return input.Ref + " == nil"
		}
	}
	return ""
}
