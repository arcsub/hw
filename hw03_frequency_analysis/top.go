package hw03frequencyanalysis

import (
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	maxWords := 10

	split := strings.Fields(text)

	words := make(map[string]int)

	for _, el := range split {
		words[el]++
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
