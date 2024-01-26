package main

// final O(n) ver
func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n

	reverse := func(start, end int) {
		for start < end {
			nums[end], nums[start] = nums[start], nums[end]
			start++
			end--
		}
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)

	return nums
}

// v1 which is O(k*n), ugh
// func rotate(nums []int, k int) []int {
// 	for k > 0 {
// 		for i := 0; i < len(nums)-1; i++ {
// 			nums[len(nums)-1], nums[i] = nums[i], nums[len(nums)-1]
// 		}
// 		k--
// 	}
// 	return nums
// }
