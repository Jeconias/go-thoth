// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: generator/template/struct.gohtml

package template

import (
	"github.com/lab259/go-my-ast-hurts"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Struct generates generator/template/struct.gohtml
func Struct(pkg *myasthurts.Package) string {
	var _b strings.Builder
	RenderStruct(&_b, pkg)
	return _b.String()
}

// RenderStruct render generator/template/struct.gohtml
func RenderStruct(_buffer io.StringWriter, pkg *myasthurts.Package) {
	_buffer.WriteString("\npackage ")
	_buffer.WriteString(gorazor.HTMLEscape(pkg.Name))
	_buffer.WriteString(" ")
	_buffer.WriteString(("\n\n"))
	for _, s := range pkg.Structs {

		_buffer.WriteString("type ")
		_buffer.WriteString(gorazor.HTMLEscape(s.Name()))
		_buffer.WriteString(" struct { ")

		_buffer.WriteString(("\n"))

		for _, f := range s.Fields {

			_buffer.WriteString("\t")
			_buffer.WriteString(gorazor.HTMLEscape(f.Name))
			_buffer.WriteString(" ")
			_buffer.WriteString(gorazor.HTMLEscape(f.RefType.Name()))
			_buffer.WriteString(" ")
			_buffer.WriteString(("\n"))
		}

		_buffer.WriteString("}")

	}

}
