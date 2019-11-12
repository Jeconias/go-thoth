package validation

import (
	"fmt"
	"io"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

// IsCIDRInput TODO
type IsCIDRInput struct {
	Field *myasthurts.Field
	Tag   myasthurts.TagParam
	Ref   string
}

// IsCIDR TODO
func IsCIDR(_buffer io.StringWriter, input *IsCIDRInput, v ...string) {
	var version string
	if len(v) > 0 {
		version = v[0]
	}
	rules.RenderCondition(
		_buffer,
		fmt.Sprintf("isCIDR%s(%s)", version, input.Ref),
		input.Field,
		input.Tag,
	)
}
