package main

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		stack: make([]*TreeNode, 0),
	}
	it.leftmostInorder(root)
	return it
}

func (bi *BSTIterator) leftmostInorder(node *TreeNode) {
	for node != nil {
		bi.stack = append(bi.stack, node)
		node = node.Left
	}
}

func (bi *BSTIterator) Next() int {
	topIndex := len(bi.stack) - 1
	next := bi.stack[topIndex]
	bi.stack = bi.stack[:topIndex]

	if next.Right != nil {
		bi.leftmostInorder(next.Right)
	}

	return next.Val
}

func (bi *BSTIterator) HasNext() bool {
	return len(bi.stack) > 0
}
