package main

import (
	"fmt"
	"stuBaseGo01/tree"
)

func main() {
	var root tree.Node
	root.SetVal(3)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = &tree.Node{2, nil, nil}
	root.Right.Left = &tree.Node{}
	root.Traversal()
	count := 0
	root.TraversalFunc(func(node *tree.Node) {
		count += 1
	})
	fmt.Println(count)
	max := 0
	cl := root.TraversalFuncForChannel()
	for node := range cl {
		if node.Val > max {
			max = node.Val
		}
	}
	fmt.Println(max)
}

//--------------------------------扩展
type myTree struct {
	node *tree.Node
}

/**
左子遍历
*/
func (myNode *myTree) myTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//myTree{myNode.node.Left}.myTraversal() 报错 找不到地址 传过来的是一个指针地址 需要一个变量接
	lNode := myTree{myNode.node.Left}
	lNode.myTraversal()
	rNode := myTree{myNode.node.Right}
	rNode.myTraversal()
	myNode.node.Print()
}
func main02() {
	var root tree.Node
	root.SetVal(3)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = &tree.Node{2, nil, nil}
	root.Right.Left = &tree.Node{}
	m := myTree{&root}
	m.myTraversal()
}
