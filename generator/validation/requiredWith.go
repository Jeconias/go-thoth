package validation

import (
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// RequiredWithInput TODO
type RequiredWithInput struct {
	Struct    *myasthurts.Struct
	StructRef string
	Field     *myasthurts.Field
	Tag       myasthurts.TagParam
	Ref       string
}

// RequiredWith TODO
func RequiredWith(_buffer io.StringWriter, input *RequiredWithInput, args ...string) {
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
	rules.RenderEvaluation(_buffer, exp, expressions, " || ", input.Field, input.Tag)
}

func requiredWith(field *myasthurts.Field, ref string) string {
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType, *myasthurts.ArrayRefType, *myasthurts.ChanRefType:
			return " Empty(len(" + ref + "))"
		case *myasthurts.StarRefType:
			return ref + " == nil"
		}
	}
	return ""
}