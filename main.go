package main

import "fmt"

var (
	nums = []int{0, 1, 2, 3, 4, 0, 0, 7, 8, 9, 10, 11, 12, 0}
	k    = 1
)

func main() {
	fmt.Print(ContainsNearbyDuplicate(nums, k))
}
