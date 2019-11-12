package validation

import (
	"fmt"
	"io"

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
func IsCompare(_buffer io.StringWriter, input *IsCompareInput, op string) {
	rules.RenderCondition(
		_buffer,
		isCompare(input, op),
		input.Field,
		input.Tag,
	)
}

func isCompare(input *IsCompareInput, op string) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`len(%s) %s %s`, input.Ref, op, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil || len(*%s) %s %s`, input.Ref, input.Ref, op, input.Value)
		}
	case "int":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`%s %s %s`, input.Ref, op, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil`, input.Ref)
		}
	}
	return ""
}
