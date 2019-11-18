package validation

import (
	"fmt"
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

func hasLengthOf(input *HasLengthOfInput) (condition string) {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf("len(%s) != %s", input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || len(*%s) != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		}
	case "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "int", "int8", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf("%s != %s", input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || len(*%s) != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf("len(%s) != %s", input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		}
	}
	return ""
}
