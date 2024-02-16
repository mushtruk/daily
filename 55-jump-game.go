package main

// mr - 2,
// []int{2, 3, 1, 1, 4}

func CanJump(nums []int) bool {
	var maxReach int

	for i, num := range nums {
		if i > maxReach {
			return false
		}

		if num+i > maxReach {
			maxReach = num + i
		}
	}
	return maxReach >= len(nums)-1
}
