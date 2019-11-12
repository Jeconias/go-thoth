package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsGteInput TODO
type IsGteInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsGte TODO
func IsGte(_buffer io.StringWriter, input *IsGteInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isGte(input),
		input.Field,
		input.Tag,
	)
}

func isGte(input *IsGteInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`len(%s) >= %s`, input.Ref, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil`, input.Ref)
		}
	}
	return ""
}
