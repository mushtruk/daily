package main

var (
	nums = []int{3, 2, 2, 3}
	val  = 3
)

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
	}
	return len(nums)
}
