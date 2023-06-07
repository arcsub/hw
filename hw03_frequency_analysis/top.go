package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	var (
		maxWords  = 10
		separator = regexp.MustCompile(`[.,?!;:\s']+`)
		tire      = regexp.MustCompile(`^[-—]+$|^$`)
	)

	loweredText := strings.ToLower(text)
	separatedText := separator.Split(loweredText, -1)

	words := make(map[string]int)

	for _, word := range separatedText {
		if tire.MatchString(word) {
			continue
		}

		words[word]++
	}

	result := maps.Keys(words)

	sort.Slice(result, func(i, j int) bool {
		if words[result[i]] == words[result[j]] {
			return result[i] < result[j]
		}

		return words[result[i]] > words[result[j]]
	})

	if len(result) > maxWords {
		result = result[:maxWords]
	}

	return result
}
