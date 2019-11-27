package validation

import (
	"fmt"
	"io"
	"strings"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// OmitemptyInput TODO
type OmitemptyInput struct {
	Struct    *myasthurts.Struct
	StructRef string
	Field     *myasthurts.Field
	Tag       myasthurts.TagParam
	Ref       string
}

// Omitempty TODO
func Omitempty(_buffer io.StringWriter, input *OmitemptyInput) {
	var expressions = make([]string, 0)
	var values []string
	var tagValue string
	for _, opt := range input.Tag.Options {
		opts := strings.Split(opt, "=")
		switch len(opts) {
		case 1:
			fmt.Println("1\t", opts[0]) // Regexp or Function
		case 2:
			key := opts[0]
			value := opts[1]

			if fn, ok := bakedInValidators[key]; ok {
				expressions = append(expressions, fn(&ValidateInput{
					Field: input.Field,
					Tag:   input.Tag,
					Ref:   input.Ref,
					Value: value,
				}))

				values = append(values, fmt.Sprintf("[%s = %v]", key, value))

				continue
			}

			panic(NewErrTypeNotSupported(key, input.Field))
		default:
			panic("omitempty")
		}
	}

	condition := fmt.Sprintf("!%s", required(input.Field, input.Ref))

	if len(expressions) > 0 {
		preCondition := fmt.Sprintf("!(%s)", strings.Join(expressions, " && "))
		condition = fmt.Sprintf("%s && %s", condition, preCondition)
		tagValue = strings.Join(values, " ~ ")
	} else {
		tagValue = "omitempty"
	}

	rules.RenderConditionWithValue(_buffer, condition, input.Field.Name, tagValue)
}
