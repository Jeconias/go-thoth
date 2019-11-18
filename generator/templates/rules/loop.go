// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/rules/loop.gohtml

package rules

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Loop generates templates/rules/loop.gohtml
func Loop(condition string, array string, field *myasthurts.Field, tag myasthurts.TagParam) string {
	var _b strings.Builder
	RenderLoop(&_b, condition, array, field, tag)
	return _b.String()
}

// RenderLoop render templates/rules/loop.gohtml
func RenderLoop(_buffer io.StringWriter, condition string, array string, field *myasthurts.Field, tag myasthurts.TagParam) {
	_buffer.WriteString("\nfor _, v := range ")
	_buffer.WriteString(gorazor.HTMLEscape(array))
	_buffer.WriteString(" {\n    if ")
	_buffer.WriteString((condition))
	_buffer.WriteString(" {\n        errs = append(errs, NewError(\"")
	_buffer.WriteString(gorazor.HTMLEscape(field.Name))
	_buffer.WriteString("\", \"")
	_buffer.WriteString(gorazor.HTMLEscape(tag.Value))
	_buffer.WriteString("\"))\n    }\n}")

}