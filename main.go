package main

import "fmt"

var (
	nums   = []int{2, 3, 4}
	target = 6
)

func main() {
	fmt.Print(TwoSum(nums, target))
}
