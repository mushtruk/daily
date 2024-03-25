package main

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	tail := merged

	current1 := list1
	current2 := list2

	for current1 != nil || current2 != nil {

		if current1 == nil {
			tail.Next = current2
			break
		}

		if current2 == nil {
			tail.Next = current1
			break
		}

		if current1.Val > current2.Val {
			tail.Next = current2
			current2 = current2.Next
		} else if current2.Val > current1.Val {
			tail.Next = current1
			current1 = current1.Next
		} else {
			tail.Next = current1
			current1 = current1.Next
			tail = tail.Next
			tail.Next = current2
			current2 = current2.Next
		}

		tail = tail.Next
	}
	return merged.Next
}
