//go:build ignore
// +build ignore

package main

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right *= nums[i]
	}

	return answer
}
