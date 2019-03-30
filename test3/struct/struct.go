package main

import "fmt"
func main() {
	var root test3.TreeNode
	root = TreeNode{value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5,nil,nil}
	root.Right.Left = new(TreeNode)
	root.Right.Right = createNode(3)

	fmt.Println(root)

	nodes := []treeNode {
		{value: 3},
		{},
		{6,nil,&root},
	}

	fmt.Println(nodes)
	root.print(15)
}
