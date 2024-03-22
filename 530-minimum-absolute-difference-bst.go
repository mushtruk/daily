package main

import "math"

// Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func GetMinimumDifference(root *TreeNode) int {
	var (
		prevNode *TreeNode
		midDiff  = math.MaxInt32
	)

	inorderTraversal(root, &prevNode, &midDiff)

	return midDiff
}

func inorderTraversal(node *TreeNode, prevNode **TreeNode, diff *int) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left, prevNode, diff)

	if *prevNode != nil {
		currDiff := node.Val - (*prevNode).Val
		if currDiff < *diff {
			*diff = currDiff
		}
	}

	*prevNode = node

	inorderTraversal(node.Right, prevNode, diff)
}
