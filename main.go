package main

import "fmt"

func main() {
	ms := NewMinStack()
	ms.Push(4)
	ms.Push(5)
	ms.Push(1)
	fmt.Print(ms.Top())
}
