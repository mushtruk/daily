package main

import "fmt"

// You are given a sorted unique integer array nums.

// A range [a,b] is the set of all integers from a to b (inclusive).

// Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

// Each range [a,b] in the list should be output as:

// "a->b" if a != b
// "a" if a == b

func SummaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	rangeStart := 0

	for i, num := range nums {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if i == rangeStart {
				result = append(result, fmt.Sprintf("%d", num))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[rangeStart], num))
			}
			rangeStart = i + 1
		}
	}
	return result
}
