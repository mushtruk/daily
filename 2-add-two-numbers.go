package main

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead = &ListNode{}
		current   = dummyHead
		carry     = 0
	)

	for l1 != nil || l2 != nil || carry > 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10

		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next
	}

	return dummyHead.Next
}

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func PrintList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}
