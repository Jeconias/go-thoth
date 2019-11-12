package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsEthereumAddressInput TODO
type IsEthereumAddressInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsEthereumAddress is the validation function for validating if the field's value is a valid ethereum address based currently only on the format
func IsEthereumAddress(_buffer io.StringWriter, input *IsEthereumAddressInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isisEthereumAddress(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
