package main

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{})
	maxLength := 0

	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	for _, num := range nums {
		if _, exists := numSet[num-1]; !exists {
			length := 1
			for nextNum := num + 1; ; nextNum++ {
				_, exists := numSet[nextNum]
				if !exists {
					break
				}
				length++
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}
