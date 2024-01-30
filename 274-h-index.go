package main

import (
	"sort"
)

// []int{3, 0, 6, 1, 5}
// SORTED: [0 1 3 5 6]

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for i, c := range citations {
		if c <= i {
			return i
		}
	}
	return len(citations)
}
