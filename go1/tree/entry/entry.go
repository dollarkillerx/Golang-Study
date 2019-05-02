package main

import (
	"fmt"
)

func main() {
	var root tree.treeNode
	fmt.Println(root) //{0 <nil> <nil>}
	root = treeNode{3, &treeNode{}, &treeNode{}}
	fmt.Println(root) //{3 0xc00000c060 0xc00000c080}
	root = treeNode{left: &treeNode{}}
	fmt.Println(root) //{0 0xc00000c060 <nil>}
	root.right = &treeNode{value: 3}
	fmt.Println(root)               //{0 0xc00000c0e0 0xc00000c120}
	root.left.right = new(treeNode) //new 返回地址
	fmt.Println(root.left.right)    //&{0 <nil> <nil>}

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000090020}]

	fmt.Println()
	root.print(5)

	root.traverse()
}
