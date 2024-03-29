package main

// Given the head of a linked list, return the list after sorting it in ascending order.

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cutPoint := findCutPoint(head)
	secondHalf := cutPoint.Next
	cutPoint.Next = nil

	leftSorted := SortList(head)
	rightSorted := SortList(secondHalf)

	return merge(leftSorted, rightSorted)
}

func findCutPoint(node *ListNode) *ListNode {
	dummy := &ListNode{Next: node}
	slow, fast := dummy, node

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}
