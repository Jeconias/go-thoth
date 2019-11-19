package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsISBNInput TODO
type IsISBNInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsISBN TODO
func IsISBN(_buffer io.StringWriter, input *IsISBNInput) {
	condition, isLoop := isISBN(input.Field, input.Ref, "isbn")
	if isLoop {
		rules.RenderLoop(
			_buffer,
			condition,
			input.Ref,
			input.Field,
			input.Tag,
		)
	} else {
		rules.RenderCondition(
			_buffer,
			condition,
			input.Field,
			input.Tag,
		)
	}
}

func isISBN(field *myasthurts.Field, ref string, rule string) (condition string, loop bool) {
	condition = fmt.Sprintf("!(isISBN10(%s) || isISBN13(%s))", ref, ref)
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			rules.MapCondition[ref] = fmt.Sprintf("isISBN10(%s) || isISBN13(%s)", ref, ref)
			return condition, false
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf("%s == nil || !(isISBN10(*%s) || isISBN13(*%s))", ref, ref, ref)
			rules.MapCondition[ref] = fmt.Sprintf("%s == nil || (isISBN10(*%s) || isISBN13(*%s))", ref, ref, ref)
			return condition, false
		case *myasthurts.ArrayRefType:
			condition = fmt.Sprintf("!(isISBN10(%s) || isISBN13(%s))", "v", "v")
			rules.MapCondition[ref] = fmt.Sprintf("isISBN10(%s) || isISBN13(%s)", "v", "v")
			return condition, true
		}
	default:
		panic(NewErrTypeNotSupported(rule, field))
	}

	return
}
