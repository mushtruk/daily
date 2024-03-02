package main

import "fmt"

func main() {
	fmt.Print(EvalRPN([]string{"2", "1", "+", "3", "*"}))
}
