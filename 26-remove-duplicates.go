package main

var _nums = []int{1, 1, 1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 6}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueCount-1] {
			nums[uniqueCount] = nums[i]
			uniqueCount++
		}
	}

	return uniqueCount
}
