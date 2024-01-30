package main

import "fmt"

var citations = []int{3, 0, 6, 1, 5}

func main() {
	fmt.Print(hIndex(citations))
}
