package validation

import "regexp"

var usernameRegex = regexp.MustCompile("^[\\w]{4,64}$")

func ValidUsername(s string) bool {
	return usernameRegex.MatchString(s)
}
