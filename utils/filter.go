package utils

import (
	snowball "github.com/kljensen/snowball/english"
	"strings"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}
func stopwordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {},
		"be": {}, "but": {}, "by": {}, "for": {}, "if": {}, "in": {},
		"into": {}, "is": {}, "it": {}, "no": {}, "not": {}, "of": {},
		"on": {}, "or": {}, "such": {}, "that": {}, "the": {}, "their": {},
		"then": {}, "there": {}, "these": {}, "they": {}, "this": {}, "to": {},
		"was": {}, "will": {}, "with": {}}

	r := make([]string, len(tokens)) //NOTE: initialize with zero length but same capacity as tokens

	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}

	return r

}
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowball.Stem(token, false)
	}
	return r
}
