package main

import "strings"

var (
	haystack = "sadbutsad"
	needle   = "sad"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
