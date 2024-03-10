package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkNodeSymmetry(root.Left, root.Right)
}

func checkNodeSymmetry(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return checkNodeSymmetry(left.Left, right.Right) && checkNodeSymmetry(left.Right, right.Left)
}
