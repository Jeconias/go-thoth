package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsBitcoinAddressInput TODO
type IsBitcoinAddressInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsBitcoinAddress is the validation function for validating if the field's value is a valid btc address
func IsBitcoinAddress(_buffer io.StringWriter, input *IsBitcoinAddressInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isBitcoinAddress(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
