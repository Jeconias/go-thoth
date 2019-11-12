package validation

import "fmt"

func regexMatch(source string, ref string, isPtr ...bool) string {
	var ptr string
	if len(isPtr) > 0 {
		ptr = "*"
	}
	return fmt.Sprintf("%s.MatchString(%s%s)", source, ptr, ref)
}
