package main

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

// Example 1:

// Input: n = 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
// Example 2:

// Input: n = 2
// Output: false

func IsHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = powSum(splitDigits(n))
	}
	return n == 1
}

func splitDigits(n int) []int {
	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func powSum(digits []int) (result int) {
	for _, digit := range digits {
		result += digit * digit
	}
	return result
}
