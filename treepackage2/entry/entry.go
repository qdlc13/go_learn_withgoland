package main

import (
	"fmt"
	"go_learn_withgoland/treepackage"
)

// 内嵌方式 Embedding
type myTreeNode struct {
	*treepackage.TreeNode //node *treepackage.TreeNode 语法糖简洁
}

// 多理解包装组合
func (myNode *myTreeNode) postorder() {
	if myNode.TreeNode == nil || myNode == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postorder()
	right := myTreeNode{myNode.Right}
	right.postorder()
	myNode.Print()
}
func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	//root := treepackage.TreeNode{Value: 3}
	root := myTreeNode{&treepackage.TreeNode{Value: 3}}
	root.Left = &treepackage.TreeNode{}
	root.Right = &treepackage.TreeNode{5, nil, nil}
	root.Right.Left = new(treepackage.TreeNode) //go特性 等价于其他语言root.right->left 无论是地址还是结构本身一律用.访问成员
	root.Left.Right = treepackage.CreatNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse() //This method is shadowed. 如果没新定义Traverse函数，这种写法还是调用下面这个函数

	root.TreeNode.Traverse() //02345

	fmt.Println()

	//多看包装
	//myRoot := myTreeNode{&root}
	root.postorder()
	fmt.Println()

}
