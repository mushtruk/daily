package main

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

// Example 1:

// Input: nums = [1,2,3,1], k = 3
// Output: true
// Example 2:

// Input: nums = [1,0,1,1], k = 1
// Output: true
// Example 3:

// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numsToIndex := make(map[int]int)

	for i, num := range nums {
		index, exists := numsToIndex[num]
		if exists && i-index <= k {
			return true
		}
		numsToIndex[num] = i
	}

	return false
}
