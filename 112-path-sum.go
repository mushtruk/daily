package main

// Given the root of a binary tree and an integer targetSum,
// return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

// A leaf is a node with no children.

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			return true
		} else {
			return false
		}
	}

	leftHasSumPath := HasPathSum(root.Left, targetSum)
	rightHasSumPath := HasPathSum(root.Right, targetSum)

	return leftHasSumPath || rightHasSumPath
}
