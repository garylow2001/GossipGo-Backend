package utils

import "strings"

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
