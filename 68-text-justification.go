package main

import (
	"strings"
)

// var (
// 	words    = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
// 	maxWidth = 16
// )

func FullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)

	rows := divideIntoRows(words, maxWidth)
	for i, row := range rows {
		result = append(result, justify(row, maxWidth, len(rows)-1 == i))
	}

	return result
}

func divideIntoRows(words []string, maxWidth int) [][]string {
	var (
		rows       [][]string
		currentRow []string
		lineLength int
	)

	for _, word := range words {
		if lineLength+len(currentRow)+len(word) > maxWidth {
			rows = append(rows, currentRow)
			currentRow = []string{}
			lineLength = 0
		}
		currentRow = append(currentRow, word)
		lineLength += len(word)
	}
	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	return rows
}

func justify(row []string, maxWidth int, isLastRow bool) string {
	if len(row) == 1 || isLastRow {
		return leftJustify(row, maxWidth)
	}

	totalLength := sumLength(row)
	spacesNeeded := maxWidth - totalLength
	spacesBetweenWords, extraSpaces := distributeSpaces(spacesNeeded, len(row)-1)

	var justifiedStr strings.Builder
	for i, word := range row {
		justifiedStr.WriteString(word)
		if i < len(row)-1 {
			justifiedStr.WriteString(strings.Repeat(" ", spacesBetweenWords))
			if extraSpaces > 0 {
				justifiedStr.WriteString(" ")
				extraSpaces--
			}
		}
	}

	return justifiedStr.String()
}

func leftJustify(row []string, maxWidth int) string {
	justifiedStr := strings.Join(row, " ")
	spacesToEnd := maxWidth - len(justifiedStr)
	return justifiedStr + strings.Repeat(" ", spacesToEnd)
}

func sumLength(words []string) int {
	total := 0
	for _, word := range words {
		total += len(word)
	}
	return total
}

func distributeSpaces(spacesNeeded, gaps int) (int, int) {
	if gaps == 0 {
		return spacesNeeded, 0
	}
	return spacesNeeded / gaps, spacesNeeded % gaps
}
