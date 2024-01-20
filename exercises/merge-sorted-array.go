package exercises

import (
	"fmt"
	"sort"
)

var (
	nums1 = []int{1, 2, 3, 0, 0, 0}
	m     = 3
	nums2 = []int{2, 5, 6}
	n     = 3
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

func Run() {
	merge(nums1, m, nums2, n)
	fmt.Print(nums1)
}
