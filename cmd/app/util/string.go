package util

import (
	"unicode"
	"unicode/utf8"
)

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
