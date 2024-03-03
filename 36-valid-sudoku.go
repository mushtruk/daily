package main

// var SudokuBoard = [][]string{
// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
// }

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		if !validateRow(board[i]) {
			return false
		}
	}

	for j := 0; j < len(board[0]); j++ {
		if !validateColumn(board, j) {
			return false
		}
	}
	for rowStart := 0; rowStart < len(board); rowStart += 3 {
		for columnStart := 0; columnStart < len(board[0]); columnStart += 3 {
			if !validateBox(board, rowStart, columnStart) {
				return false
			}
		}
	}
	return true
}

func validateRow(row []byte) bool {
	seen := make(map[byte]struct{})
	for _, char := range row {
		if char == '.' {
			continue
		}
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}
	return true
}

func validateColumn(board [][]byte, columnIndex int) bool {
	seen := make(map[byte]struct{})

	for rowIndex := 0; rowIndex < len(board); rowIndex++ {
		cell := board[rowIndex][columnIndex]

		if cell == '.' {
			continue
		}

		if _, exists := seen[cell]; exists {
			return false
		}

		seen[cell] = struct{}{}
	}
	return true
}

func validateBox(board [][]byte, rowStart, colStart int) bool {
	seen := make(map[byte]struct{})
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := board[rowStart+i][colStart+j]

			if cell == '.' {
				continue
			}

			if _, exists := seen[cell]; exists {
				return false
			}

			seen[cell] = struct{}{}
		}
	}
	return true
}

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Example 1:

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
// Example 2:

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
