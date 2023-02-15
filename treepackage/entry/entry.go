package main

import (
	"fmt"
	"go_learn_withgoland/treepackage"
)

// 扩充系统类型或者别人的类型方法1定义别名2使用组合
// 组合方式
type myTreeNode struct {
	node *treepackage.TreeNode
}

// 多理解包装组合
func (myNode *myTreeNode) postorder() {
	if myNode.node == nil || myNode == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postorder()
	right := myTreeNode{myNode.node.Right}
	right.postorder()
	myNode.node.Print()
}

func main() {
	root := treepackage.TreeNode{3, nil, nil}
	root.Left = &treepackage.TreeNode{}
	root.Right = &treepackage.TreeNode{5, nil, nil}
	root.Right.Left = new(treepackage.TreeNode) //go特性 等价于其他语言root.right->left 无论是地址还是结构本身一律用.访问成员
	root.Left.Right = treepackage.CreatNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()

	//多看包装
	myRoot := myTreeNode{&root}
	myRoot.postorder()

}
