package validation

import "fmt"

func regexMatch(source string, ref string) string {
	return fmt.Sprintf("%s.MatchString(%s)", source, ref)
}
