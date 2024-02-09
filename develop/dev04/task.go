package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятка", "тяпка", "тяпка", "алла", "лала", "алал", "алал", "ведро"}
	res := solve(words)
	fmt.Println(res)
}

func solve(words []string) map[string]*[]string {
	m := getMap(words)
	res := make(map[string]*[]string, 0)
	for key, value := range m {
		if len(value) != 1 {
			slice := make([]string, len(value))
			i := 0
			for word, _ := range value {
				slice[i] = word
				i++
			}
			sort.Strings(slice)
			res[key] = &slice
		}
	}
	return res
}

func getMap(words []string) map[string]map[string]struct{} {
	m := make(map[string]map[string]struct{})
	for _, word := range words {
		var hasAnagram bool
		lowerWord := strings.ToLower(word)

		for key, _ := range m {
			if isAnagram(lowerWord, key) {
				m[key][lowerWord] = struct{}{}
				hasAnagram = true
			}
		}
		if !hasAnagram {
			m[lowerWord] = make(map[string]struct{})
			m[lowerWord][lowerWord] = struct{}{}
		}
	}
	return m
}

func isAnagram(word1, word2 string) bool {
	if word1 == word2 || len(word1) != len(word2) {
		return false
	}
	return getSortWord(word1) == getSortWord(word2)
}

func getSortWord(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
