package templates

import (
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

func filterValidate(_buffer io.StringWriter, field *myasthurts.Field, tag myasthurts.TagParam, value interface{}, args ...string) {
	switch tag.Value {
	case "required":
		rules.RenderRequired(_buffer, field, tag, value)
	case "eq":
		rules.RenderEq(_buffer, field, tag, value, args)
	default:
		argSplit := strings.Split(tag.Value, "=")
		tag.Value = argSplit[0]
		args := argSplit[1:]
		filterValidate(_buffer, field, tag, value, args...)
	}
}
