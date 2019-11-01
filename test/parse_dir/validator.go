package any

import "github.com/lab259/errors/v2"

var (
	errRequiredString = errors.New("string could not be empty")
	errRequiredInt64  = errors.New("int64 could not be zero")
)

func required(t string, v interface{}) bool {
	switch t {
	case "string":
		return requiredString(v.(string))
	case "int":
		return requiredInt(v.(int))
	}
	return false
}

func requiredString(v string) bool {
	return v == ""
}

func requiredInt(v int) bool {
	return v == 0
}

func requiredInt64(v int64) bool {
	return v == 0
}
