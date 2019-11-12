package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsEmailInput TODO
type IsEmailInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsEmail TODO
func IsEmail(_buffer io.StringWriter, input *IsEmailInput) {
	rules.RenderCondition(
		_buffer,
		isEmail(input),
		input.Field,
		input.Tag,
	)
}

func isEmail(input *IsEmailInput) (condition string) {
	condition = fmt.Sprintf("!%s", regexMatch("emailRegex", input.Ref))
	switch input.Field.RefType.Name() {
	case "string":
		switch input.Field.RefType.(type) {
		case *myasthurts.BaseRefType:
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || ! %s`, input.Ref, regexMatch("emailRegex", input.Ref, true))
			rules.MapCondition[input.Ref] = condition
			return condition
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			panic("not implemented")
		}
	default:
		panic("not implemented")
	}

	return
}
