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

func filterValidate(_buffer io.StringWriter, str *myasthurts.Struct, field *myasthurts.Field, tag myasthurts.TagParam, attribute interface{}, args ...string) {
	switch tag.Value {
	case "-":
		// Skip field...
	case "required":
		rules.RenderRequired(_buffer, field, tag, attribute)
	case "required_with":
		for _, s := range args {
			for _, f := range str.Fields {
				if f.Name == s {
					for _, param := range f.Tag.Params {
						rules.RenderRequired_with(_buffer, field, param, attribute)
					}
				}
			}
		}
	case "eq":
		rules.RenderEq(_buffer, field, tag, attribute, args)
	default:
		k, v := splitArgs(tag)
		tag.Value = k
		filterValidate(_buffer, str, field, tag, attribute, v...)
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
