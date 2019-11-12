package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsLtInput TODO
type IsLtInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsLt TODO
func IsLt(_buffer io.StringWriter, input *IsLtInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isLt(input),
		input.Field,
		input.Tag,
	)
}

func isLt(input *IsLtInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`len(%s) < %s`, input.Ref, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil`, input.Ref)
		}
	}
	return ""
}
