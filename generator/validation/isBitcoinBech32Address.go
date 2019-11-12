package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsBitcoinBech32AddressInput TODO
type IsBitcoinBech32AddressInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsBitcoinBech32Address TODO
func IsBitcoinBech32Address(_buffer io.StringWriter, input *IsBitcoinBech32AddressInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isBitcoinBech32Address(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
