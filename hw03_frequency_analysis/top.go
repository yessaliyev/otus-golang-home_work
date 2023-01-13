package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if text == "" {
		return make([]string, 0)
	}

	s := strings.FieldsFunc(text, split)

	if len(s) < 1 {
		return make([]string, 0)
	}

	data := map[string]int{}

	//считаем количество слов (слово: количество)
	for _, word := range s {
		if _, ok := data[word]; ok {
			data[word] += 1
			continue
		}

		data[word] = 1
	}

	topWords := getTopWords(data)
	return sortTopWords(topWords)
}

/**
Получаем топ 10 слов по количеству.
Получаем map, для дальнейшей сортировки.
*/
func getTopWords(words map[string]int) map[string]int {
	keys := make([]string, 0, len(words))
	topWords := map[string]int{}

	for key := range words {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})

	to := 10

	if len(keys) < 10 {
		to = 10 - (10 - len(keys))
	}

	for _, k := range keys[:to] {
		topWords[k] = words[k]
	}

	return topWords
}

/**
Сортируем слова по количеству и лексикографический. Группирую по val
*/
func sortTopWords(words map[string]int) []string {
	keys := make([]int, 0, len(words))
	groupedByVal := map[int][]string{}
	result := make([]string, 0, len(words))

	//группируем слова по количеству(5 => ["ты", "что"])
	//чтобы можно было отсортировать лексикографическом порядке
	for key, val := range words {
		if _, ok := groupedByVal[val]; ok {
			groupedByVal[val] = append(groupedByVal[val], key)
			continue
		}

		groupedByVal[val] = []string{key}
	}

	//сортируем лексикографическом порядке
	for _, value := range groupedByVal {
		sort.Strings(value)
	}

	//далее сортируем по ключам в порядке убывания
	for key := range groupedByVal {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, val := range keys {
		result = append(result, groupedByVal[val]...)
	}

	return result
}

func split(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
