package main

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).

// A valid BST is defined as follows:

// The left
// subtree
//  of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

func IsValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, low, high *int) bool {
	if node == nil {
		return true
	}

	if (low != nil && node.Val <= *low) || (high != nil && node.Val >= *high) {
		return false
	}

	return validate(node.Left, low, &node.Val) && validate(node.Right, &node.Val, high)
}
