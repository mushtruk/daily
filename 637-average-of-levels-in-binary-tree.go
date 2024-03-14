package main

// Given the root of a binary tree, return the average value of the nodes on each level in the form of an array.

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	queue := []*TreeNode{root}
	averages := []float64{}

	for len(queue) > 0 {
		var (
			lvlSize = len(queue)
			lvlSum  = 0.0
		)

		for i := 0; i < lvlSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			lvlSum += float64(currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		averages = append(averages, lvlSum/float64(lvlSize))
	}

	return averages
}
