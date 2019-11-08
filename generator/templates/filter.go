package templates

import (
	"fmt"
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

func filterValidate(_buffer io.StringWriter, field *myasthurts.Field, tag myasthurts.TagParam, value interface{}, args ...string) {
	switch tag.Value {
	case "-":
		// Skip field...
	case "required":
		rules.RenderRequired(_buffer, field, tag, value)
	case "eq":
		rules.RenderEq(_buffer, field, tag, value, args)
	default:
		k, v := splitArgs(tag)
		tag.Value = k
		filterValidate(_buffer, field, tag, value, v...)
	}
}

func splitArgs(tag myasthurts.TagParam) (key string, value []string) {
	args := strings.Split(tag.Value, "=")
	if len(args) > 0 {
		key = args[0]
		value = args[1:]
		if len(value) == 2 {
			v := strings.Split(args[1], " ")
			fmt.Println("Args\t", len(args), args)
			fmt.Println("Values\t", len(v), v)
			return key, v
		}
	}
	return
}
