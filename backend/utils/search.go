package utils

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func FuzzySearch(text, query string) []string {
	lines := strings.Split(text, "\n")
	var results []string
	for _, line := range lines {
		if fuzzy.MatchFold(query, line) {
			results = append(results, line)
		}
	}
	return results
}
