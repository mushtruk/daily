package exercises

import (
	"sort"
)

var (
	nums1 = []int{1, 2, 3, 0, 0, 0}
	m     = 3
	nums2 = []int{2, 5, 6}
	n     = 3
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var result = make([]int, 0, m+n)
	result = append(result, nums1[:m]...)
	result = append(result, nums2[:n]...)
	sort.Ints(result)
	copy(nums1, result)
}
