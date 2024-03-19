package main

// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// SearchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 7)

func SearchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	left, right := 0, r*c-1

	for left <= right {
		mid := left + (right-left)/2
		row := mid / r
		col := mid % c

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
