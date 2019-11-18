package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsHexcolorInput TODO
type IsHexcolorInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsHexcolor TODO
func IsHexcolor(_buffer io.StringWriter, input *IsHexcolorInput) {
	rules.RenderCondition(
		_buffer,
		isRegex("hexcolorRegex", input.Field, input.Ref, "hexcolor"),
		input.Field,
		input.Tag,
	)
}

func isRegex(regexp string, field *myasthurts.Field, ref string, rule string) (condition string) {
	condition = fmt.Sprintf("!%s", regexMatch(regexp, ref))
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			rules.MapCondition[ref] = condition
			return condition
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || ! %s`, ref, regexMatch(regexp, ref, true))
			rules.MapCondition[ref] = condition
			return condition
		case *myasthurts.ArrayRefType:
			condition = fmt.Sprintf("!%s", regexMatch(regexp, "v"))
			rules.MapCondition[ref] = condition
			return condition
		}
	default:
		panic(NewErrTypeNotSupported(rule, field))
	}

	return
}
