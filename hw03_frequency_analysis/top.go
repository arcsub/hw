package hw03frequencyanalysis

import (
	"regexp"
	"sort"

	"golang.org/x/exp/maps"
)

var RE = regexp.MustCompile(`\s+`)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	split := RE.Split(text, -1)

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

	return result[:10]
}
