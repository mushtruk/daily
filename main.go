package main

import "fmt"

var (
	ransomNote = "aa"
	magazine   = "ab"
)

func main() {
	fmt.Print(CanConstruct(ransomNote, magazine))
}
