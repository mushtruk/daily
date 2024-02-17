package main

import (
	"sort"
)

// var nums = []int{-1, 0, 1, 2, -1, -4}

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for fixed := 0; fixed < len(nums)-2; fixed++ {
		if fixed > 0 && nums[fixed] == nums[fixed-1] {
			continue
		}
		left, right := fixed+1, len(nums)-1

		for left < right {
			sum := nums[fixed] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[fixed], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
