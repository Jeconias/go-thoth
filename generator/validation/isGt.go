package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsGtInput TODO
type IsGtInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsGt TODO
func IsGt(_buffer io.StringWriter, input *IsGtInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isGt(input),
		input.Field,
		input.Tag,
	)
}

func isGt(input *IsGtInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`len(%s) > %s`, input.Ref, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil`, input.Ref)
		}
	}
	return ""
}
