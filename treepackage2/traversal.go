package treepackage2

import (
	"fmt"
)

// 中序遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 改进中序遍历
func (node *TreeNode) Traverse2() {
	node.TraverseFunc(func(n *TreeNode) {
		n.Print()
	})
	fmt.Println()
}

func (node *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// 使用channel中序遍历
func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node //为啥没锁住？？？
		})
		close(out) //关闭的channel仍然可以读取数据
	}()
	return out
}
