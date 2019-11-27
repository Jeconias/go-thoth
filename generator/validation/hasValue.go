package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// RequiredInput TODO
type RequiredInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// HasValue TODO
func HasValue(_buffer io.StringWriter, input *RequiredInput) {
	rules.RenderRequired(_buffer, &rules.RequiredInput{
		Field: input.Field,
		Tag:   input.Tag,
		Ref:   input.Ref,
	})
}

func required(field *myasthurts.Field, ref string) (condition string) {
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf(`Empty(len(%s))`, ref)
			rules.MapCondition[ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil`, ref)
			rules.MapCondition[ref] = condition
			return
		}
	case "int":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf("IsInt(%s)", ref)
			rules.MapCondition[ref] = fmt.Sprintf("! IsInt(%s)", ref)
			return
		case *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf(`Empty(len(%s))`, ref)
			rules.MapCondition[ref] = condition
			return
		}
	case "bool":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			condition = fmt.Sprintf("! %s", ref)
			rules.MapCondition[ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf("! %s == nil || * %s", ref, ref)
			rules.MapCondition[ref] = condition
			return
		}
	default:
		panic(NewErrTypeNotSupported("required", field))
	}

	return
}
