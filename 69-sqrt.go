package main

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.

func MySqrt(x int) int {

	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2
		squareX := mid * mid
		if squareX == x {
			return mid
		}
		if squareX > x {
			high = mid - 1
		}
		if squareX < x {
			low = mid + 1
		}
	}

	return high
}
