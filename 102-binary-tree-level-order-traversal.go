package main

// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func LevelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)

	if root == nil {
		return levels
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvlSize := len(queue)
		lvl := []int{}
		for i := 0; i < lvlSize; i++ {

			currentNode := queue[0]
			queue = queue[1:]

			lvl = append(lvl, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		levels = append(levels, lvl)
	}

	return levels
}
