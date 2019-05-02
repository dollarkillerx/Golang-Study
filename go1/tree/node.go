package tree

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	// 工厂函数控制struct的创建
	return &treeNode{value: value}
	// 重点: 局部遍历也可以返回给外面用
	// 创建在堆栈不需要知道
	// 如果局部变量取地址返回使用就会分配在堆上  参与垃圾回收
	// 反之没有使用就会分配到栈上
}

// (node treeNode)为接受者 为结构定义方法
func (node *treeNode) print(a int) {
	node.value = 5
	fmt.Println(node, " ", a)
}
