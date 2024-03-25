package main

func main() {
	l1 := CreateList([]int{1, 2, 4})
	l2 := CreateList([]int{1, 3, 4})

	mergedList := MergeTwoLists(l1, l2)

	PrintList(mergedList)
}
