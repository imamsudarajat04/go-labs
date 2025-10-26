package utils

import (
	"strings"
	"unicode"

	"github.com/gertd/go-pluralize"
)

var plural = pluralize.NewClient()

func Capitalize(subject string) string {
	if subject == "" {
		return ""
	}

	runes := []rune(subject)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Pluralize(word string) string {
	return plural.Plural(strings.ToLower(word))
}
