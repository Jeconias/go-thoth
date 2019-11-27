package validation

import myasthurts "github.com/lab259/go-my-ast-hurts"

var bakedInValidators = map[string]func(input *ValidateInput) string{
	"min": hasMinOf,
	"max": hasMaxOf,
}

// ValidateInput TODO
type ValidateInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
	Value interface{}
}

func hasMinOf(input *ValidateInput) string {
	return isCompare(input, "<=", "min")
}

func hasMaxOf(input *ValidateInput) string {
	return isCompare(input, ">=", "max")
}
