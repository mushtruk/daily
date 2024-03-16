package main

// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
// (i.e., from left to right, then right to left for the next level and alternate between).

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func ZigzagLevelOrder(root *TreeNode) [][]int {
	zigZag := make([][]int, 0)

	if root == nil {
		return zigZag
	}

	queue := []*TreeNode{root}
	direction := true

	for len(queue) > 0 {
		lvlSize := len(queue)
		lvl := make([]int, lvlSize)

		for i := 0; i < lvlSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			insertPos := i
			if !direction {
				insertPos = lvlSize - 1 - i
			}

			lvl[insertPos] = currentNode.Val

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

		}
		zigZag = append(zigZag, lvl)
		direction = !direction
	}
	return zigZag
}
