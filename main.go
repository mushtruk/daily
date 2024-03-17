package main

import "fmt"

func main() {
	tree := ArrayToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	PrintTree(tree, 0, "root")
	fmt.Print(Serialize(tree))
}
