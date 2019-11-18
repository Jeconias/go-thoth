package validation

import (
	"fmt"
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// RequiredWithAllInput TODO
type RequiredWithAllInput struct {
	Struct    *myasthurts.Struct
	StructRef string
	Field     *myasthurts.Field
	Tag       myasthurts.TagParam
	Ref       string
}

// RequiredWithAll TODO
func RequiredWithAll(_buffer io.StringWriter, input *RequiredWithAllInput, args ...string) {
	var expressions = make([]string, 0)
	for _, s := range args {
		for _, f := range input.Struct.Fields {
			if f.Name == s {
				ref := input.StructRef + "." + f.Name
				for range f.Tag.Params {
					expressions = append(expressions, rules.MapCondition[ref])
				}
			}
		}
	}

	exp := requiredWith(input.Field, input.Ref)
	preCondition := fmt.Sprintf("(%s)", strings.Join(expressions, " && "))
	condition := fmt.Sprintf("%s && %s", preCondition, exp)
	rules.RenderCondition(_buffer, condition, input.Field, input.Tag)
}

func requiredWithAll(field *myasthurts.Field, ref string) (condition string) {
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			condition = fmt.Sprintf("Empty(len(%s))", ref)
			rules.MapCondition[ref] = condition
			return
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf("%s == nil", ref)
			rules.MapCondition[ref] = condition
			return
		}
	default:
		panic(NewErrTypeNotSupported("required_with_all", field))
	}

	return
}
