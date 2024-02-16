package main

import "strings"

func Every(slice []string, fn func(string) bool) bool {
	for _, value := range slice {
		if !fn(value) {
			return false
		}
	}
	return true
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for prefix != "" {
		if Every(strs, func(word string) bool {
			return strings.HasPrefix(word, prefix)
		}) {
			return prefix
		}
		prefix = prefix[:len(prefix)-1]
	}

	return prefix
}
