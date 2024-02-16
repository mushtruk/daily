package main

import (
	"strings"
	"unicode"
)

// var sp = "A man, a plan, a canal: Panama"

// v2
func IsPalindrome(s string) bool {
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

// v1

// func isPalindrome(s string) bool {
// 	str := strings.ToLower(removeNonAlphanumeric(s))
// 	reversedStr := new(strings.Builder)

// 	for i := len(str) - 1; i >= 0; i-- {
// 		reversedStr.WriteByte(str[i])
// 	}

// 	return str == reversedStr.String()
// }

// func removeNonAlphanumeric(s string) string {
// 	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
// 	if err != nil {
// 		panic(err)
// 	}

// 	return reg.ReplaceAllString(s, "")
// }
