package validation

import (
	"fmt"
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsCompareInput TODO
type IsCompareInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsCompare TODO
func IsCompare(_buffer io.StringWriter, input *IsCompareInput, op string, rule string) {
	rules.RenderCondition(
		_buffer,
		isCompare(input, op, rule),
		input.Field,
		input.Tag,
	)
}

func isCompare(input *IsCompareInput, op string, rule string) (condition string) {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf(`len(%s) %s %s`, input.Ref, op, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || len(*%s) %s %s`, input.Ref, input.Ref, op, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		}
	case "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "int", "int8", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			return fmt.Sprintf(`%s %s %s`, input.Ref, op, input.Value)
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || len(*%s) != %s`, input.Ref, input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf("len(%s) != %s", input.Ref, input.Value)
			rules.MapCondition[input.Ref] = condition
			return
		}
	default:
		if strings.HasPrefix(input.Field.RefType.Name(), "map") {
			switch input.Field.RefType.(type) {
			case *myasthurts.BaseRefType:
				condition = fmt.Sprintf(`len(%s) %s %s`, input.Ref, op, input.Value)
				rules.MapCondition[input.Ref] = condition
				return
			case *myasthurts.StarRefType:
				condition = fmt.Sprintf(`%s == nil || len(*%s) %s %s`, input.Ref, input.Ref, op, input.Value)
				rules.MapCondition[input.Ref] = condition
				return
			}
		} else {
			switch input.Field.RefType.(type) {
			case *myasthurts.BaseRefType:
				if input.Field.RefType.Type() != nil {
					condition = fmt.Sprintf("%s == %s{}", input.Ref, input.Field.RefType.Type().Name())
					rules.MapCondition[input.Ref] = condition
					return
				}

				condition = fmt.Sprintf("IsValid(%s)", input.Ref)
				rules.MapCondition[input.Ref] = condition
				return

			case *myasthurts.StarRefType:
				condition = fmt.Sprintf(`%s == nil || len(*%s) %s %s`, input.Ref, input.Ref, op, input.Value)
				rules.MapCondition[input.Ref] = condition
				return
			case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
				condition = fmt.Sprintf(`len(%s) %s %s`, input.Ref, op, input.Value)
				rules.MapCondition[input.Ref] = condition
				return
			default:
				panic(NewErrTypeNotSupported(rule, input.Field))
			}
		}
	}
	return
}
