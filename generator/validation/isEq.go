package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsEqInput TODO
type IsEqInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsEq TODO
func IsEq(_buffer io.StringWriter, input *IsEqInput) {
	rules.RenderCondition(_buffer,
		isEq(input),
		input.Field,
		input.Tag,
	)
}

func isEq(input *IsEqInput) (condition string) {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf(`%s != "%s"`, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || *%s != "%s"`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			panic("not implemented")
		}
	case "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "int", "int8", "int16", "int32", "int64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf(`%s != %s`, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || *%s != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		}
	case "float32", "float64":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf(`%s != %s`, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || *%s != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		}
	case "complex64", "complex128":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf(`%s != %s`, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || *%s != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return condition
		}
	default:
		panic("not implemented")
	}

	return
}
