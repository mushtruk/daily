package main

import "fmt"

var (
	s = "anagram"
	t = "nagaram"
)

func main() {
	fmt.Print(IsAnagram(s, t))
}
