package main

// []int{2, 3, 1, 1, 4}

func Jump(nums []int) int {
	var jumps, maxReach, nextMaxReach int

	for i, num := range nums {
		if i > maxReach {
			break
		}

		if num+i > nextMaxReach {
			nextMaxReach = num + i
		}

		if i == maxReach {
			if i < len(nums)-1 {
				jumps++
			}
			maxReach = nextMaxReach
		}
	}

	return jumps
}
