package main

import "fmt"

func IsNumberPalindrome(x int) bool {
	digits := fmt.Sprint(x)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
