package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsLteInput TODO
type IsLteInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsLte TODO
func IsLte(_buffer io.StringWriter, input *IsLteInput, args ...string) {
	rules.RenderCondition(
		_buffer,
		isLte(input),
		input.Field,
		input.Tag,
	)
}

func isLte(input *IsLteInput) string {
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return fmt.Sprintf(`len(%s) <= %s`, input.Ref, input.Value.(string))
		case *myasthurts.StarRefType:
			return fmt.Sprintf(`%s == nil`, input.Ref)
		}
	}
	return ""
}
