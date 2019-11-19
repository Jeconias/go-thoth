package validation

import (
	"fmt"

	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/lab259/go-thoth/generator/templates/rules"
)

func isRegex(regexp string, field *myasthurts.Field, ref string, rule string) (condition string, loop bool) {
	condition = fmt.Sprintf("!%s", regexMatch(regexp, ref))
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			rules.MapCondition[ref] = condition
			return condition, false
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || ! %s`, ref, regexMatch(regexp, ref, true))
			rules.MapCondition[ref] = condition
			return condition, false
		case *myasthurts.ArrayRefType:
			condition = fmt.Sprintf("!%s", regexMatch(regexp, "v"))
			rules.MapCondition[ref] = condition
			return condition, true
		}
	default:
		panic(NewErrTypeNotSupported(rule, field))
	}

	return
}

func isFunc(regexp string, field *myasthurts.Field, ref string, rule string) (condition string, loop bool) {
	condition = fmt.Sprintf("!%s(%s)", regexp, ref)
	switch field.RefType.Name() {
	case "string":
		switch field.RefType.(type) {
		case *myasthurts.BaseRefType:
			rules.MapCondition[ref] = fmt.Sprintf("%s(%s)", regexp, ref)
			return condition, false
		case *myasthurts.StarRefType:
			condition = fmt.Sprintf(`%s == nil || !%s(*%s)`, ref, regexp, ref)
			rules.MapCondition[ref] = fmt.Sprintf(`%s == nil || %s(*%s)`, ref, regexp, ref)
			return condition, false
		case *myasthurts.ArrayRefType:
			condition = fmt.Sprintf("!%s(%s)", regexp, "v")
			rules.MapCondition[ref] = fmt.Sprintf("%s(%s)", regexp, "v")
			return condition, true
		}
	default:
		panic(NewErrTypeNotSupported(rule, field))
	}

	return
}
