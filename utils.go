package main

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	sanitizeNameRegex, _ = regexp.Compile("[^a-zA-Z0-9:_]")
	mergeUScoreRegex, _  = regexp.Compile("__+")
)

func safeName(dirty string) string {
	return mergeUScoreRegex.ReplaceAllString(
		sanitizeNameRegex.ReplaceAllString(
			strings.ToLower(dirty), "_"),
		"_")
}

func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
