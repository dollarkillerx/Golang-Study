package tree

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print(1)
	node.right.traverse()
	node.print(2)
}
