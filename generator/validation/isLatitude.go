package validation

import (
	"fmt"
	"io"
	"strconv"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsLatitudeInput TODO
type IsLatitudeInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

// IsLatitude TODO
func IsLatitude(_buffer io.StringWriter, input *IsLatitudeInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isLatitude(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}

// TODO (@edumarcal): To parser gohtml
func isLatitude(field *myasthurts.Field, ref string) bool {
	var v string
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			v = ref
		}
	case "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			v = strconv.FormatUint(2, 10)
		}
	case "int", "int8", "int16", "int32", "int64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			v = strconv.FormatInt(2, 10)
		}
	case "float32":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			v = strconv.FormatFloat(0, 'f', -1, 32)
		}
	case "float64":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			v = strconv.FormatFloat(0, 'f', -1, 64)
		}
	}
	return latitudeRegex.MatchString(v)
}
