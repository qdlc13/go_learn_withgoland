package treepackage

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 没有构造函数但可以有工厂函数
func CreatNode(value int) *TreeNode {
	return &TreeNode{Value: value} //c++中是错误返回了一个局部变量的地址 go中可用

}

func (node TreeNode) Print() { //方法前的参数（使用者）也是传值类型
	fmt.Print(node.Value) //值传递把node的内容拷贝一份给print函数打印
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value //等价(*node).value = value
}
