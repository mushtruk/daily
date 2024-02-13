package main

import (
	"strings"
	"unicode"
)

var sp = "A man, a plan, a canal: Panama"

func isPalindrome(s string) bool {
	var b strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			b.WriteRune(unicode.ToLower(c))
		}
	}

	str := b.String()
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
