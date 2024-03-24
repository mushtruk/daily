package main

func main() {
	l1 := CreateList([]int{2, 4, 3})
	l2 := CreateList([]int{5, 6, 4})

	result := AddTwoNumbers(l1, l2)

	PrintList(result)
}
