package main

// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	visible := []int{}

	for len(queue) > 0 {

		lvlSize := len(queue)

		for i := 0; i < lvlSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if i == lvlSize-1 {
				visible = append(visible, currentNode.Val)
			}

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

	}
	return visible
}
