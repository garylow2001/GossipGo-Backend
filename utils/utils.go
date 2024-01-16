package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SlugToString(s string) string {
	if len(s) == 0 {
		return s
	}
	s = strings.ReplaceAll(s, "-", " ")
	s = cases.Title(language.AmericanEnglish).String(s)

	return s
}
