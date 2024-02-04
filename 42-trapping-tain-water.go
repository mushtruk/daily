package main

var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

func trap(height []int) int {
	var (
		leftMax  = make([]int, len(height))
		rightMax = make([]int, len(height))
	)

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	water := 0
	for i := 0; i < len(height); i++ {
		water += max(0, min(leftMax[i], rightMax[i])-height[i])
	}

	return water
}
