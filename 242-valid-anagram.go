package main

import "unicode/utf8"

func IsAnagram(s string, t string) bool {
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
