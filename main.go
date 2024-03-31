package main

import "fmt"

func main() {
	// valid
	tree1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// invalid
	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println("Tree 1 is a valid BST:", IsValidBST(tree1)) // 1
	fmt.Println("Tree 1 is a valid BST:", IsValidBST(tree2)) // 0
}
