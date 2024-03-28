package main

// Given an integer array nums where the elements are sorted in ascending order, convert it to a
// height-balanced
//
//	binary search tree.

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return constructBST(nums, 0, len(nums)-1)
}

func constructBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = constructBST(nums, left, mid-1)
	root.Right = constructBST(nums, mid+1, right)

	return root
}
