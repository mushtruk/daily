package main

import (
	"strings"
)

// var regularString = "the sky is blue"

func ReverseWords(s string) string {
	words := strings.Fields(s)
	reversed := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}

	return strings.Join(reversed, " ")
}
