package validation

import (
	"fmt"
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// Dive TODO
func Dive(_buffer io.StringWriter, input *ValidateInput) {
	rules.RenderLoop(
		_buffer,
		dive(input.Field, input.Ref),
		input.Ref,
		input.Field,
		input.Tag,
	)
}

func dive(field *myasthurts.Field, ref string) (condition string) {
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.ArrayRefType:
			condition = "Empty(len(v))"
			rules.MapCondition[ref] = condition
			return
		default:
			panic(NewErrTypeNotSupported("dive", field))
		}
	case "uint":
		switch field.RefType.(type) {
		case *myasthurts.ArrayRefType:
			condition = fmt.Sprintf(`Empty(len(%s))`, ref)
			rules.MapCondition[ref] = condition
			return
		default:
			panic(NewErrTypeNotSupported("dive", field))
		}
	case "int":
		switch field.RefType.(type) {
		case *myasthurts.ArrayRefType:
			condition = "IsInt(v)"
			rules.MapCondition[ref] = condition
			return
		default:
			panic(NewErrTypeNotSupported("dive", field))
		}
	default:
		if strings.HasPrefix(field.RefType.Name(), "map") {
			switch field.RefType.(type) {
			case *myasthurts.BaseRefType:
				condition = "Empty(len(v))"
				rules.MapCondition[ref] = condition
				return
			default:
				panic(NewErrTypeNotSupported("dive", field))
			}
		}

		panic(NewErrTypeNotSupported("dive", field))

	}
}
