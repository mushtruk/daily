package main

import "strings"

// var str = "   fly me   to   the moon  "

func LengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
