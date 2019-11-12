package templates

import (
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/validation"
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
	Struct    *myasthurts.Struct
	StructRef string
	Field     *myasthurts.Field
	Tag       myasthurts.TagParam
	Ref       string
}

func filterValidate(_buffer io.StringWriter, input *FilterInput, args ...string) {
	switch input.Tag.Value {
	case "-":
		// Skip field...
	case "required":
		validation.HasValue(_buffer, &validation.RequiredInput{
			Field: input.Field,
			Tag:   input.Tag,
			Ref:   input.Ref,
		})
	case "required_with":
		validation.RequiredWith(_buffer, &validation.RequiredWithInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_with_all":
		validation.RequiredWithAll(_buffer, &validation.RequiredWithAllInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_without":
		validation.RequiredWithout(_buffer, &validation.RequiredWithoutInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "required_without_all":
		validation.RequiredWithoutAll(_buffer, &validation.RequiredWithoutAllInput{
			Struct:    input.Struct,
			StructRef: input.StructRef,
			Field:     input.Field,
			Tag:       input.Tag,
			Ref:       input.Ref,
		}, args...)
	case "len":
		if len(args) == 1 {
			validation.HasLengthOf(_buffer, &validation.HasLengthOfInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	case "eq":
		if len(args) == 1 {
			validation.IsEq(_buffer, &validation.IsEqInput{
				Field: input.Field,
				Tag:   input.Tag,
				Ref:   input.Ref,
				Value: args[0],
			})
		}
	default:
		// TODO: (@edumarcal) to fix not match rule
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
