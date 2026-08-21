package wordfreq

import (
	"strings"
	"unicode"
)

// WordFrequency takes a string and returns a map of each word (lowercased,
// punctuation stripped) to the number of times it appears in the input.
func WordFrequency(s string) map[string]int {
	freq := make(map[string]int)

	// Split on whitespace to get raw tokens.
	tokens := strings.Fields(s)

	for _, token := range tokens {
		// Strip leading and trailing punctuation from each token.
		word := strings.TrimFunc(token, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})

		// Skip tokens that are entirely punctuation/whitespace.
		if word == "" {
			continue
		}

		// Normalise to lowercase for case-insensitive counting.
		freq[strings.ToLower(word)]++
	}

	return freq
}
