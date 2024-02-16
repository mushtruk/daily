package main

func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var (
		rows       = make([]string, numRows)
		currentRow = 0
		goingDown  = false
		result     = ""
	)

	for _, char := range s {
		rows[currentRow] += string(char)

		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currentRow += 1
		} else {
			currentRow -= 1
		}
	}

	for _, row := range rows {
		result += row
	}

	return result
}
