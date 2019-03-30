package test3

import "fmt"

type TreeNode struct {
	Value int
	Left,Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

// 接受者 (接受者) 函数名称
func (node TreeNode) Print(a int) {
	fmt.Println(node,a)
}