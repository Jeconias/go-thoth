package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsFQDNInput TODO
type IsFQDNInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsFQDN TODO
func IsFQDN(_buffer io.StringWriter, input *IsFQDNInput) {
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isFQDN(%s)", input.Ref),
		input.Field,
		input.Tag,
	)
}
