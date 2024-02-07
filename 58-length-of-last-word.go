package main

import "strings"

var str = "   fly me   to   the moon  "

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
