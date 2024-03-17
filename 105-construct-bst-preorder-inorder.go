package main

import (
	"fmt"
	"strconv"
)

// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
// and inorder is the inorder traversal of the same tree, construct and return the binary tree.

//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	i := 0
	for ; inorder[i] != rootVal; i++ {
	}

	root.Left = BuildTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree(preorder[i+1:], inorder[i+1:])

	return root
}

func Serialize(root *TreeNode) []string {
	result := []string{}

	serializeHelper(root, &result)

	return result
}

func Deserialize(tree []string) *TreeNode {
	var index int

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(tree) {
			return nil
		}

		if tree[index] == "#" {
			index++
			return nil
		}

		val, err := strconv.Atoi(tree[index])
		if err != nil {
			return nil
		}
		node := &TreeNode{Val: val}

		index++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	return dfs()
}

func serializeHelper(node *TreeNode, result *[]string) {
	if node == nil {
		*result = append(*result, "#")
		return
	}
	*result = append(*result, strconv.Itoa(node.Val))
	serializeHelper(node.Left, result)
	serializeHelper(node.Right, result)
}

func ArrayToTreeNode(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}

	i := 1
	for len(queue) > 0 && i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if arr[i] != nil {
			current.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

func PrintTree(node *TreeNode, level int, prefix string) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "       "
	}
	format += "---" + prefix + ": "

	fmt.Println(format + strconv.Itoa(node.Val))
	PrintTree(node.Left, level+1, "L")
	PrintTree(node.Right, level+1, "R")
}
