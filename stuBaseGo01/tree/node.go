package tree

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func (n *Node) Print() {
	fmt.Println(n.Val)
}
func (n *Node) SetVal(val int) {
	if n == nil {
		return
	}
	n.Val = val
}

/**
中序变量树
*/
func (n *Node) Traversal() {
	if n == nil {
		return
	}
	n.Print()
	n.Left.Traversal()
	n.Right.Traversal()
}

func (n *Node) TraversalFunc(f func(node *Node)) {
	if n == nil {
		return
	}
	f(n)
	n.Left.TraversalFunc(f)
	n.Right.TraversalFunc(f)
}

func (n *Node) TraversalFuncForChannel() chan *Node {
	cl := make(chan *Node)
	go func() {
		n.TraversalFunc(func(no *Node) {
			cl <- no
		})
		close(cl)
	}()
	return cl
}
