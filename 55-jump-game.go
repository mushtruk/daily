package main

// mr - 2, 
// []int{2, 3, 1, 1, 4}

func canJump(nums []int) bool {
	var maxReach int

	for i, num := range nums {
		if i > maxReach {
			return false
		}

		if num+i > maxReach {
			maxReach = num + i
		}
	}
	if maxReach >= len(nums)-1 {
		return true
	}

	return false
}
