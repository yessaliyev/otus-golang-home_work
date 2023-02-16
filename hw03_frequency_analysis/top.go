package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if text == "" {
		return nil
	}

	s := strings.FieldsFunc(text, split)

	if len(s) < 1 {
		return nil
	}

	data := map[string]int{}

	// считаем количество слов (слово: количество)
	for _, word := range s {
		data[word]++
	}

	return getTopWords(data)
}

func getTopWords(words map[string]int) []string {
	keys := make([]string, 0, len(words))

	for key := range words {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})

	to := len(keys)

	if to > 10 {
		to = 10
	}

	return keys[:to]
}

func split(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
