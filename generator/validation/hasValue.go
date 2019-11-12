package validation

import (
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
