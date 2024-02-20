package main

import "fmt"

var (
	pattern = "abba"
	s       = "dog cat cat dog"
)

func main() {
	fmt.Print(WordPattern(pattern, s))
}
