package templates

import (
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

func hasTag(structsThoth []*myasthurts.Struct) bool {
	for _, s := range structsThoth {
		for _, field := range s.Fields {
			if len(field.Tag.Params) > 0 {
				return true
			}
		}
	}
	return false
}

// FilterInput TODO
type FilterInput struct {
	Struct       *myasthurts.Struct
	StructRef    string
	Field        *myasthurts.Field
	Tag          myasthurts.TagParam
	AttributeRef string
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

func filterValidate(_buffer io.StringWriter, input *FilterInput, args ...string) {
	switch input.Tag.Value {
	case "-":
		// Skip field...
	case "required":
		rules.RenderRequired(_buffer, &rules.RequiredInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.AttributeRef,
		})
	case "required_with":
		var expressions = append(make([]string, 0), requiredWith(input.Field, input.AttributeRef))
		for _, s := range args {
			for _, f := range input.Struct.Fields {
				if f.Name == s {
					ref := input.StructRef + "." + f.Name
					for range f.Tag.Params {
						expressions = append(expressions, rules.Condition[ref])
					}

				}
			}
		}
		condition := strings.Join(expressions, " || ")
		rules.RenderEvaluation(_buffer, condition, input.Field, input.Tag)
	case "eq":
		if len(args) == 1 {
			rules.RenderEq(_buffer, &rules.EqInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.AttributeRef,
				Value: args[0],
			})
		}
	default:
		k, v := splitArgs(input.Tag)
		input.Tag.Value = k
		filterValidate(_buffer, input, v...)
	}
}

func splitArgs(tag myasthurts.TagParam) (string, []string) {
	args := strings.Split(tag.Value, "=")
	switch len(args) {
	case 1:
		return args[0], args[1:]
	case 2:
		v := strings.Split(args[1], " ")
		return args[0], v
	default:
		return "", nil
	}
}
