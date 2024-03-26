package main

// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

// [5,7,7,8,8,10]

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	result[0] = findOccurance(nums, target, true)
	if result[0] != -1 {
		result[1] = findOccurance(nums, target, false)
	}
	return result
}

func findOccurance(nums []int, target int, findFirst bool) int {
	left, right, answer := 0, len(nums)-1, -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			answer = mid
			if findFirst {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return answer
}
