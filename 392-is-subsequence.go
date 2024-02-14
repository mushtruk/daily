package main

var (
	isSubsequenceInput  = "abc"
	isSubsequenceTarget = "ahbgdc"
)

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sPos := 0

	for tPos := 0; tPos < len(t); tPos++ {
		if s[sPos] == t[tPos] {
			sPos++
			if sPos == len(s) {
				return true
			}
		}
	}

	return false
}
