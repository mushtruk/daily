package main

import "strings"

// Given a pattern and a string s, find if s follows the same pattern.

// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

// Example 1:

// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Example 2:

// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false
// Example 3:

// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false

func WordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	patternMap := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if mappedWord, exists := patternMap[char]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			for _, w := range patternMap {
				if w == word {
					return false // word already mapped to a different char
				}
			}
			patternMap[char] = word
		}
	}

	return true
}
