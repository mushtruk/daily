package main

func removeDuplicates2(nums []int) int {
	var occurrences = 1
	if len(nums) <= occurrences {
		return len(nums)
	}

	uniqueCount := occurrences
	for i := occurrences; i < len(nums); i++ {
		if nums[i] != nums[uniqueCount-occurrences] {
			nums[uniqueCount] = nums[i]
			uniqueCount++
		}
	}
	return uniqueCount
}
