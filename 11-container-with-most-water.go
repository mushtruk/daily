package main

// var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

func MaxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0

	for left < right {
		currentArea := min(height[left], height[right]) * (right - left)

		maxArea = max(currentArea, maxArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
