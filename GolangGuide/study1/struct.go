package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 接受者 (接受者) 函数名称
func (node treeNode) print(a int) {
	fmt.Println(node,a)
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.right.right = createNode(3)

	fmt.Println(root)

	nodes := []treeNode {
		{value: 3},
		{},
		{6,nil,&root},
	}

	fmt.Println(nodes)
	root.print(15)
}