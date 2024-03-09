package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
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
		"a":    {},
		"an":   {},
		"the":  {},
		"and":  {},
		"be":   {},
		"have": {},
		"i":    {},
		"in":   {},
		"of":   {},
		"at":   {},
		"that": {},
		"to":   {},
	}
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		stemmedToken := snowballeng.Stem(token, false)
		r = append(r, stemmedToken)
	}
	return r
}
