package main

import "fmt"

func main() {

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	fmt.Println("Minimum Absolute Difference:", GetMinimumDifference(root))
}
