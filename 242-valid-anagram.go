package main

import (
	"sort"
	"unicode/utf8"
)

func IsStringAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sortedS := SortString(s)
	sortedT := SortString(t)
	return sortedS == sortedT
}

func SortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	return string(chars)
}

func IsUnicodeAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range t {
		r, _ := utf8.DecodeRuneInString(string(char))
		charCount[r]--
		if charCount[r] < 0 {
			return false
		}
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}
