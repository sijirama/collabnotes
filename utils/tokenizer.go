package utils

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) //NOTE: return true for the delimiter function if the char is not a letter or a number
	})
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens) // change to lower case
	tokens = stopwordFilter(tokens)  // remove common words
	tokens = stemmerFilter(tokens)   // stem to root
	return tokens
}
